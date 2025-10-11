package main

import "fmt"

func main() {
	var num1, num2 float64
	var operator string
	fmt.Println("Enter First Number")
	fmt.Scan(&num1)
	fmt.Println("Enter Operator:(-,+,*,/)")
	fmt.Scan(&operator)
	fmt.Println("Enter First Number")
	fmt.Scan(&num2)

	switch operator {
	case "-":
		fmt.Println("Result:", num1-num2)
	case "+":
		fmt.Println("Result:", num1+num2)
	case "*":
		fmt.Println("Result:", num1*num2)
	case "/":
		fmt.Println("Result:", num1/num2)
	default:
		fmt.Println("enter some valued operator")
	}
}
