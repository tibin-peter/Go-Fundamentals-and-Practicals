package main

import "fmt"

func printEven(ch chan bool, done chan bool) {
	for i := 2; i <= 10; i += 2 {
		<-ch // receive chanel waiting for signal from odd
		fmt.Println("Even:", i)
		done <- true // pass signal back
	}

}

func prilntOdd(ch chan bool, done chan bool) {
	for i := 1; i <= 9; i += 2 {
		fmt.Println("Odd:", i)
		ch <- true // singnal even goroutine pass
		<-done     // wait for evn finish  receive
	}
}

func main() {

	ch := make(chan bool)
	done := make(chan bool)

	go printEven(ch, done)
	prilntOdd(ch, done) // starts with this function odd

}
