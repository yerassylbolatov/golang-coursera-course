package main

import "fmt"

func main() {
	cancelCh := make(chan struct{})
	writeCh := make(chan int)

	go func(cancel chan struct{}, write chan int) {
		val := 0
		for {
			select {
			case <-cancel:
				return
			case write <- val:
				val++
			}
		}
	}(cancelCh, writeCh)

	for i := range writeCh {
		fmt.Println("writeCh val", i)
		if i > 3 {
			fmt.Println("writeCh closed")
			cancelCh <- struct{}{}
			break
		}
	}
}
