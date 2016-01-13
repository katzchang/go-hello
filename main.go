package main

import (
	"fmt"
	"html"
	"net/http"
	"encoding/json"
)

func main() {
	http.HandleFunc("/foo", handler)
	http.ListenAndServe(":8081", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("a") == "hoge" {
		http.Error(w, "kuke", http.StatusBadRequest)
	}
	fmt.Fprintf(w, "Hello, %q\n", html.EscapeString(r.URL.Path))
	fmt.Fprintln(w, "%q", hoge())
}

type Message struct {
	Name string `json:"name"`
	BodyHoge string `json:"body_hoge"`
	Time int64 `json:"time"`
}

func hoge() string {
	m := Message{"Alice", "Hel", 1294706395881547000}
	b, err := json.Marshal(m)
	fmt.Println(err)
	return string(b)
}
