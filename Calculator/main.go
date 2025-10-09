package main

import "fmt"

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("Enter the First Number:")
	fmt.Scan(&num1)
	fmt.Println("Enter the First operator(+,-,/,*):")
	fmt.Scan(&operator)
	fmt.Println("Enter the Second Number:")
	fmt.Scan(&num2)

	switch operator {
	case "+":
		fmt.Println("Result:", num1+num2)
	case "-":
		fmt.Println("Result:", num1-num2)
	case "*":
		fmt.Println("Result:", num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("cant divide by zero")
		} else {
			fmt.Println("Result:", num1/num2)
		}
	default:
		fmt.Println("Invalid operator. Please use +, -, *, or /.")

	}

}
