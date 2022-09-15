package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	ID     int
	Name   string
	Active bool
}

func main() {
	tmpl, err := template.
		New("").
		ParseFiles("method.html")
	if err != nil {
		panic(err)
	}

	users := []User{
		{1, "yerassyl", true},
		{2, "bolatov", true},
		{3, "talgatovich", false},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "method.html",
			struct {
				Users []User
			}{
				users,
			})
		if err != nil {
			panic(err)
		}
	})
	fmt.Println("server starts at :8080")
	http.ListenAndServe(":8080", nil)
}

func (u *User) PrintActive() string {
	if !u.Active {
		return ""
	}
	return "method says user " + u.Name + " is active!"
}
