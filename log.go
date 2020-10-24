package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

func LogMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := rand.Intn(10000)
		t := time.Now()
		log.Printf("Request n%d started: %s %s", id, r.Method, r.URL.Path)
		handler.ServeHTTP(w, r)
		log.Printf("Request n%d completed: %s %s. Time: %s", id, r.Method, r.URL.Path, time.Now().Sub(t))
	})
}

