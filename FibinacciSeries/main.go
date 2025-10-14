package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter how many terms: ")
	fmt.Scan(&n)

	a, b := 0, 1

	fmt.Print("Fibonacci Series: ")

	for i := 0; i < n; i++ {
		fmt.Printf("%d", a)
		next := a + b
		a = b
		b = next
	}
}
