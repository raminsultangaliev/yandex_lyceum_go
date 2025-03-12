package main

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Access Granted")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Please log in")
	cookie := http.Cookie {
		Name: "user_id",
		Value: "123",
	}
	http.SetCookie(w, &cookie)
}

func cookieMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("session")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func main() {
	http.HandleFunc("/", cookieMiddleware(Handler))
	http.HandleFunc("/login", LoginHandler)
	http.ListenAndServe("localhost:8080", nil)
}