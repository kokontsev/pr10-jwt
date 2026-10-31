package auth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type User struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Hash  []byte `json:"-"`
}
type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	Kind  string `json:"kind"`
	jwt.RegisteredClaims
}
type Server struct {
	secret  []byte
	ttl     time.Duration
	users   map[string]User
	mu      sync.Mutex
	revoked map[string]time.Time
}
type key struct{}

func New(secret string, ttl time.Duration) (*Server, error) {
	if len(secret) < 16 || ttl <= 0 {
		return nil, errors.New("JWT_SECRET must contain at least 16 bytes; TTL must be positive")
	}
	s := &Server{secret: []byte(secret), ttl: ttl, users: map[string]User{}, revoked: map[string]time.Time{}}
	for i, role := range []string{"admin", "user"} {
		h, e := bcrypt.GenerateFromPassword([]byte("secret123"), 12)
		if e != nil {
			return nil, e
		}
		email := role + "@example.com"
		s.users[email] = User{int64(i + 1), email, role, h}
	}
	return s, nil
}
func write(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
func fail(w http.ResponseWriter, status int, e string) {
	write(w, status, map[string]string{"error": e})
}
func (s *Server) sign(u User, kind string, ttl time.Duration) (string, error) {
	b := make([]byte, 16)
	if _, e := rand.Read(b); e != nil {
		return "", e
	}
	now := time.Now()
	c := Claims{u.Email, u.Role, kind, jwt.RegisteredClaims{Subject: strconv.FormatInt(u.ID, 10), Issuer: "pr10", Audience: jwt.ClaimStrings{"pr10-clients"}, IssuedAt: jwt.NewNumericDate(now), ExpiresAt: jwt.NewNumericDate(now.Add(ttl)), ID: hex.EncodeToString(b)}}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(s.secret)
}
func (s *Server) parse(raw, kind string) (*Claims, error) {
	c := &Claims{}
	t, e := jwt.ParseWithClaims(raw, c, func(t *jwt.Token) (any, error) { return s.secret, nil }, jwt.WithValidMethods([]string{"HS256"}), jwt.WithIssuer("pr10"), jwt.WithAudience("pr10-clients"), jwt.WithExpirationRequired())
	if e != nil || !t.Valid || c.Kind != kind || c.ID == "" {
		return nil, errors.New("invalid_token")
	}
	return c, nil
}
func (s *Server) pair(u User) (map[string]string, error) {
	a, e := s.sign(u, "access", s.ttl)
	if e != nil {
		return nil, e
	}
	r, e := s.sign(u, "refresh", 7*24*time.Hour)
	return map[string]string{"access": a, "token": a, "refresh": r}, e
}
func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	var in struct{ Email, Password string }
	if json.NewDecoder(r.Body).Decode(&in) != nil || in.Email == "" || in.Password == "" {
		fail(w, 400, "invalid_json")
		return
	}
	u, ok := s.users[strings.ToLower(strings.TrimSpace(in.Email))]
	if !ok || bcrypt.CompareHashAndPassword(u.Hash, []byte(in.Password)) != nil {
		fail(w, 401, "invalid_credentials")
		return
	}
	p, e := s.pair(u)
	if e != nil {
		fail(w, 500, "token_error")
		return
	}
	write(w, 200, p)
}
func (s *Server) Refresh(w http.ResponseWriter, r *http.Request) {
	var in struct{ Refresh string }
	if json.NewDecoder(r.Body).Decode(&in) != nil {
		fail(w, 400, "invalid_json")
		return
	}
	c, e := s.parse(in.Refresh, "refresh")
	if e != nil {
		fail(w, 401, "invalid_refresh")
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now()
	for id, exp := range s.revoked {
		if !exp.After(now) {
			delete(s.revoked, id)
		}
	}
	if _, ok := s.revoked[c.ID]; ok {
		fail(w, 401, "refresh_revoked")
		return
	}
	u, ok := s.users[c.Email]
	if !ok {
		fail(w, 401, "invalid_refresh")
		return
	}
	p, e := s.pair(u)
	if e != nil {
		fail(w, 500, "token_error")
		return
	}
	s.revoked[c.ID] = c.ExpiresAt.Time
	write(w, 200, p)
}
func (s *Server) AuthN(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := strings.Fields(r.Header.Get("Authorization"))
		if len(p) != 2 || !strings.EqualFold(p[0], "Bearer") {
			fail(w, 401, "unauthorized")
			return
		}
		c, e := s.parse(p[1], "access")
		if e != nil {
			fail(w, 401, "unauthorized")
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), key{}, c)))
	})
}
func roles(allowed ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, _ := r.Context().Value(key{}).(*Claims)
			if c != nil {
				for _, role := range allowed {
					if c.Role == role {
						next.ServeHTTP(w, r)
						return
					}
				}
			}
			fail(w, 403, "forbidden")
		})
	}
}
func (s *Server) Router() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("POST /api/v1/login", s.Login)
	m.HandleFunc("POST /api/v1/refresh", s.Refresh)
	m.Handle("GET /api/v1/me", s.AuthN(roles("admin", "user")(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := r.Context().Value(key{}).(*Claims)
		write(w, 200, map[string]string{"id": c.Subject, "email": c.Email, "role": c.Role})
	}))))
	m.Handle("GET /api/v1/admin/stats", s.AuthN(roles("admin")(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { write(w, 200, map[string]int{"users": len(s.users)}) }))))
	m.Handle("GET /api/v1/users/{id}", s.AuthN(roles("admin", "user")(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := r.Context().Value(key{}).(*Claims)
		id := r.PathValue("id")
		if c.Role != "admin" && id != c.Subject {
			fail(w, 403, "forbidden")
			return
		}
		for _, u := range s.users {
			if strconv.FormatInt(u.ID, 10) == id {
				write(w, 200, u)
				return
			}
		}
		fail(w, 404, "not_found")
	}))))
	return m
}
