package main

import "fmt"

func main() {
	var num int // declare variable
	fmt.Println("Enter a Number")
	fmt.Scan(&num) // take input from user

	if num%2 == 0 {
		fmt.Println(num, "is Even")
	} else {
		fmt.Println(num, "is Odd")
	}
}
