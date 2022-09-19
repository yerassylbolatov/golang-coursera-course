package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
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
	result := ""
	for i := 0; i < 1000; i++ {
		post := Post{ID: i, Text: "New Text", Time: time.Now()}
		rawJson, _ := json.Marshal(post)
		result += string(rawJson)
	}
	time.Sleep(3 * time.Millisecond)
	w.Write([]byte(result))
}

func main() {
	runtime.GOMAXPROCS(4)
	http.HandleFunc("/", handle)
	fmt.Println("server starts at :8080")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
