package main

import "fmt"

func calculate(a, b float64) (sum, diff, product, quotient float64) {
	sum = a + b
	diff = a - b
	product = a * b
	quotient = a / b
	return
}

func main() {
	s, d, p, q := calculate(20, 5)
	fmt.Println("Sum:", s)
	fmt.Println("Difference:", d)
	fmt.Println("Product:", p)
	fmt.Println("Quotient:", q)
}
