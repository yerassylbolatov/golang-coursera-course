package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Id     int
	Name   string
	Active bool
}

func main() {
	tmpl := template.Must(template.ParseFiles("users.html"))

	users := []User{
		{1, "yerassyl", true},
		{2, "<i>bolatov</i>", true},
		{3, "talgatovich", false},
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w,
			struct {
				Users []User
			}{
				users,
			})
	})
	fmt.Println("server starts at :8080")
	http.ListenAndServe(":8080", nil)
}
