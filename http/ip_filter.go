package main

import (
	"fmt"
	"net/http"
)

func Handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Access Granted")
}

func ipBlockerMiddleware(blockedIP string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := r.Header.Get("X-Real-IP")
		if ip == blockedIP {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func main() {
	http.HandleFunc("/", ipBlockerMiddleware("192.168.0.1", Handler1))
	http.ListenAndServe("localhost:8080", nil)
}