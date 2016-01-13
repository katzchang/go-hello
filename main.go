package main

import (
	"fmt"
//	"html"
	"net/url"
	"net/http"
	"encoding/json"
)

func main() {
	http.HandleFunc("/foo", foo)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8081", nil)
}

type Message struct {
	Name string `json:"n"`
	Body string `json:"b"`
}

func hoge(s string) string {
	m := Message{"Alice", s}
	b, _ := json.Marshal(m)
	return string(b)
}

func foo(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("a") == "hoge" {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, string(hoge("error!!!!")), )
	} else {
		//	fmt.Fprintf(w, "Hello, %q\n", html.EscapeString(r.URL.Path))
		fmt.Fprintln(w, url.QueryEscape(hoge("hoge-")))
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	a := query.Get("a")
	var m Message
	if err := json.Unmarshal([]byte(a), &m); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, string(hoge("error!!!!")), )
	} else {
		b, _ := json.Marshal(m)
		fmt.Fprintln(w, string(b))
	}

}
