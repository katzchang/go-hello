package main

import (
	"fmt"
	"html"
	"net/http"
)

func main() {
	http.HandleFunc("/foo", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("a") == "hoge" {
		http.Error(w, "kuke", http.StatusBadRequest)
	}
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
