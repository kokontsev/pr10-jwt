package auth

import (
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestTokenValidation(t *testing.T) {
	s, e := New("local-study-secret-change-me", time.Minute)
	if e != nil {
		t.Fatal(e)
	}
	u := s.users["user@example.com"]
	raw, e := s.sign(u, "access", -time.Minute)
	if e != nil {
		t.Fatal(e)
	}
	if _, e = s.parse(raw, "access"); e == nil {
		t.Fatal("expired accepted")
	}
	raw, e = s.sign(u, "refresh", time.Minute)
	if e != nil {
		t.Fatal(e)
	}
	if _, e = s.parse(raw, "access"); e == nil {
		t.Fatal("refresh accepted as access")
	}
	raw, e = s.sign(u, "access", time.Minute)
	if e != nil {
		t.Fatal(e)
	}
	s.secret = []byte("another-secret-at-least-16")
	if _, e = s.parse(raw, "access"); e == nil {
		t.Fatal("invalid signature accepted")
	}
}
func TestRefreshSingleUseConcurrent(t *testing.T) {
	s, e := New("local-study-secret-change-me", time.Minute)
	if e != nil {
		t.Fatal(e)
	}
	p, e := s.pair(s.users["user@example.com"])
	if e != nil {
		t.Fatal(e)
	}
	var wg sync.WaitGroup
	codes := make(chan int, 2)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			r := httptest.NewRequest("POST", "/api/v1/refresh", strings.NewReader(`{"refresh":"`+p["refresh"]+`"}`))
			w := httptest.NewRecorder()
			s.Router().ServeHTTP(w, r)
			codes <- w.Code
		}()
	}
	wg.Wait()
	close(codes)
	counts := map[int]int{}
	for c := range codes {
		counts[c]++
	}
	if counts[200] != 1 || counts[401] != 1 {
		t.Fatalf("single use failed: %v", counts)
	}
}
