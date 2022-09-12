package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "URL:", r.URL.String())
	if err != nil {
		panic(err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Starting server at :8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Print(err)
	}
}
