package main

import (
	"fmt"
	"net/http"
	"time"
	"encoding/json"
)

type Response struct {
	Greetings string `json:"greetings"`
	Name      string `json:"name"`
}

type RPCResponse struct {
	Status string   `json:"status"`
	Result interface{} `json:"result"`
}

func RPC(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				response := RPCResponse {
					Status: "error",
					Result: map[string]interface{}{},
				}
				json.NewEncoder(w).Encode(response)
			}
		}()
		next.ServeHTTP(w, r)
	}
}

func Sanitize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if !isLatin(name) {
			panic("Invalid name")
		}
		next.ServeHTTP(w, r)
	}
}

func SetDefaultName(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			r.URL.RawQuery = "name=stranger"
		}
		next.ServeHTTP(w, r)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	response := RPCResponse {
		Status: "ok",
		Result: Response {
			Greetings: "hello",
			Name:      name,
		},
	}
	json.NewEncoder(w).Encode(response)
}

var firstNumber, secondNumber = 0, 1
var requestCount = 0

func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%d", firstNumber)
	firstNumber, secondNumber = secondNumber, firstNumber + secondNumber
	requestCount += 1
}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "rpc_duration_milliseconds_count %d", requestCount)
}

func Metrics(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r* http.Request) {
		next.ServeHTTP(w, r)
	}
}

func Authorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func AnswerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "The answer is 42")
}

func StartServer(t time.Duration) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", RPC(Sanitize(SetDefaultName(HelloHandler))))
	// mux.HandleFunc("/", FibonacciHandler)
	// mux.HandleFunc("/metrics", Metrics(MetricsHandler))
	// mux.HandleFunc("/answer", Authorization(AnswerHandler))
	http.ListenAndServe(":8080", mux)
}

// func main() {
// 	StartServer(time.Second)
// }