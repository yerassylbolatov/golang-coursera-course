package main

import "fmt"

func main() {
	in := make(chan int)
	go func(out chan<- int) {
		for i := 0; i <= 4; i++ {
			fmt.Println("Before:", i)
			out <- i
			fmt.Println("After", i)
		}
		close(out)
		fmt.Println("generator is finished")
	}(in)
	for i := range in {
		fmt.Println("\tget:", i)
	}
}

//func main() {
//	ch1 := make(chan int, 1)
//
//	go func(in chan int) {
//		val := <-in
//		fmt.Println("GO: from chan", val)
//		fmt.Println("GO: after read from chan")
//	}(ch1)
//
//	ch1 <- 42
//	ch1 <- 30
//	fmt.Println("MAIN: form main go routine, after put to chan")
//	fmt.Scanln()
//}
