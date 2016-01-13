package main

import (
	"fmt"
//	"html"
	"net/http"
	"encoding/json"
)

func main() {
	http.HandleFunc("/foo", handler)
	http.ListenAndServe(":8081", nil)
}

type Message struct {
	Name string `json:"name"`
	BodyHoge string `json:"body_hoge"`
	Time int64 `json:"time"`
}

func hoge(s string) string {
	m := Message{"Alice", s, 1294706395881547000}
	b, _ := json.Marshal(m)
	return string(b)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("a") == "hoge" {
		w.Header().Set("Content-Type","application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, string(hoge("error!!!!")), )
	} else {
		//	fmt.Fprintf(w, "Hello, %q\n", html.EscapeString(r.URL.Path))
		fmt.Fprintln(w, hoge("hoge-"))
	}
}
