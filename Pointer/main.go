package main

import "fmt"

func Modified(num *int) {
	*num = *num * 2
}

func main() {
	x := 10

	Modified(&x)
	fmt.Println("New value:", x)

}
