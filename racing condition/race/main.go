package main

import (
	"fmt"
	"time"
)

var increasing int32

func inc() {
	increasing++
}

func main() {
	for i := 0; i < 1000; i++ {
		go inc()
	}
	time.Sleep(1)
	fmt.Println(increasing)
}
