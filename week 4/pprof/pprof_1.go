package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

type Post struct {
	ID       int
	Text     string
	Author   string
	Comments string
	Time     time.Time
}

func handle(w http.ResponseWriter, r *http.Request) {
	s := ""
	for i := 0; i < 1000; i++ {
		p := &Post{ID: i, Text: "New Post"}
		s += fmt.Sprintf("%#v", p)
	}
	w.Write([]byte(s))
}

func main() {
	http.HandleFunc("/", handle)
	fmt.Println("server starts at :8080")
	http.ListenAndServe(":8080", nil)
}
