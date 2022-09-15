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

//
//func (u *User) isActive() string {
//	if !u.Active {
//		return ""
//	}
//	return "User " + u.Name + " is active!"
//}

func IsUserOdd(u *User) bool {
	return u.ID%2 != 0
}

func main() {
	tmplFuncs := template.FuncMap{
		"OddUser": IsUserOdd,
	}
	tmpl, err := template.New("").Funcs(tmplFuncs).ParseFiles("func.html")
	if err != nil {
		panic(err)
	}
	users := []User{
		{1, "yerassyl", true},
		{2, "bolatov", true},
		{3, "talgatovich", false},
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "func.html", struct {
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
