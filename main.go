package main

import (
	"fmt"
	"time"
)

func Even(ch chan int) {
	for i := 0; i <= 100; i += 2 {
		<-ch
		fmt.Println("number:", i)
		ch <- 1
	}
}

func Odd(ch chan int) {
	for i := 1; i <= 100; i += 2 {
		<-ch
		fmt.Println("number:", i)
		ch <- 1
	}
}

func main() {
	ch := make(chan int, 1)
	ch <- 1
	go Even(ch)
	go Odd(ch)
	time.Sleep(time.Second)
}
