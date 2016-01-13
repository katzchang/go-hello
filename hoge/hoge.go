package hoge

import (
	"encoding/json"
)

type Message struct {
	Name string `json:"n"`
	Body string `json:"b"`
}

func Foo(s string) Message {
	return Message{"Alice", s}
}

func Bar(s string) Message {
	var m Message
	if err := json.Unmarshal([]byte(s), &m); err == nil {
		return Message{"error!!!", s}
	} else if b, e := json.Marshal(m); e == nil {
		return Message{"error!!!!!!!!!!!!!!!", s}
	} else {
		return b
	}
}
