package main

import (
	"custompackage/arithmetic" // replace 'project' with your module name
	"fmt"
)

func main() {
	a, b := 12.0, 4.0

	fmt.Println("Sum:", arithmetic.Add(a, b))
	fmt.Println("Difference:", arithmetic.Subtract(a, b))
	fmt.Println("Product:", arithmetic.Multiply(a, b))

	quotient, err := arithmetic.Divide(a, b)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Quotient:", quotient)
	}
}
