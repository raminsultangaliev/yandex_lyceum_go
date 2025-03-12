package main

import (
	"fmt"
	"net/http"
)

func languageHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("lang")
	if err != nil {
		cookie = &http.Cookie{
			Name:  "lang",
			Value: "en",
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}

	message := "Hello!" 
	if cookie.Value == "ru" {
		message = "Привет!"
	}

	fmt.Fprintf(w, message)
}

func main() {
	http.HandleFunc("/", languageHandler)
	http.ListenAndServe("localhost:8080", nil)
}	