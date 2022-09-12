package main

import (
	"fmt"
	"net/http"
)

func handler1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("RequestID", "da1241324123")
	fmt.Fprintln(w, "Your browser is", r.UserAgent())
	fmt.Fprintln(w, "You accept", r.Header.Get("Accept"))
}

func main() {
	http.HandleFunc("/", handler1)
	fmt.Println("server starts at :8080")
	http.ListenAndServe(":8080", nil)
}
