package main

import (
	"log"
	"net/http"
)

// middleware is a function that takes a HandlerFunc and returns a HandlerFunc
// it is being used to set CORS headers
func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, json")

		// If this is a preflight request, respond immediately
		if r.Method == "OPTIONS" {
			w.Header().Set("Content-Type", "application/json")
			log.Println("processed preflight request")
			return
		}

		next.ServeHTTP(w, r)
	})
}
