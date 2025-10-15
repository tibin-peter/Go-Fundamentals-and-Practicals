package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("producing:", i)
		ch <- i //send i into channel
		time.Sleep(time.Second)
	}
	close(ch) //closing chanel when done
}

func consumer(ch chan int) {
	for item := range ch { // receive untill close
		fmt.Println("consuming:", item)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	ch := make(chan int) // creating channel

	go producer(ch)
	consumer(ch)
}
