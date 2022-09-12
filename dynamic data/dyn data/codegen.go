package main

import "fmt"

type User2 struct {
	Id       int
	RealName string `unpack:"-"`
	Login    string
	Flags    int
}

func main() {
	data := []byte{
		128, 36, 17, 0,

		9, 0, 0, 0,
		118, 46, 114, 111, 109, 97, 110, 111, 118,

		16, 0, 0, 0,
	}

	u := User2{}
	u.Unpack(data)
	fmt.Printf("Unpacked user %v", u)
}
