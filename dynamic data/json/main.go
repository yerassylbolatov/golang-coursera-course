package main

import (
	"encoding/json"
	"fmt"
)

var jsonS = `[
	{"id": "42", "username": "yerassyl", "phone": "123"},
	{"id": 42, "username": "yerassyl", "phone": 123}
]`

func main() {
	var user1 interface{}
	json.Unmarshal([]byte(jsonS), &user1)
	fmt.Printf("%#v\n\n", user1)

	user2 := map[string]interface{}{
		"id":       42,
		"username": "yerassyl",
	}
	result, _ := json.Marshal(user2)
	fmt.Printf("json string: %s\n", string(result))
}
