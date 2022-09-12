package main

import (
	"fmt"
	"net/http"
)

func servers(addr string) {
	fmt.Printf("starting server at %s\n", addr)
	mux := http.NewServeMux()
	mux.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Addr:", addr, "\nURL:", r.URL.String())
		})
	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}
	server.ListenAndServe()
}

func main() {
	go servers(":8081")
	servers(":8080")
}
