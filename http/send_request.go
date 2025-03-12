package main

import (
	"fmt"
	"net/http"
	"io"
)

func startServer(address string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from server")
	})
	http.ListenAndServe(address, nil)
}

func sendRequest(url string) (string, error) {
	resp, err := http.Get("http://" + url)
	if err != nil {
		return "", fmt.Errorf("sendRequest: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	return string(body), nil
}