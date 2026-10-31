package main

import (
	"example.com/pr10/internal/auth"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ttl := 15 * time.Minute
	if v := os.Getenv("JWT_TTL"); v != "" {
		var e error
		ttl, e = time.ParseDuration(v)
		if e != nil {
			log.Fatal(e)
		}
	}
	s, e := auth.New(os.Getenv("JWT_SECRET"), ttl)
	if e != nil {
		log.Fatal(e)
	}
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal((&http.Server{Addr: ":" + port, Handler: s.Router(), ReadHeaderTimeout: 5 * time.Second}).ListenAndServe())
}
