package main

import (
	"fmt"
	"time"
)

func main() {
	//timerUse()
	//tickerUse()
	afterFunc()
}

func sayHello() {
	fmt.Println("Hello World!")
}
func afterFunc() {
	timer := time.AfterFunc(time.Second, sayHello)
	fmt.Scanln()
	timer.Stop()
	fmt.Scanln()
}

func tickerUse() {
	ticker := time.NewTicker(time.Second)
	i := 0
	for tickerTime := range ticker.C {
		i++
		fmt.Println("Step,", i, "\tTime:", tickerTime)
		if i > 3 {
			fmt.Println("tickerTime > 3 seconds, program is stopped")
			ticker.Stop()
			break
		}
	}
}

func timerUse() {
	timer := time.NewTimer(time.Second)
	i := 0
	for {
		select {
		case <-timer.C:
			fmt.Println("Time is out")
			return
		default:
			fmt.Println("Iteration:", i)
			i++
		}
	}
}
