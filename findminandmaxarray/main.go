package main

import "fmt"

func main() {
	arr := [6]int{23, 56, 12, 99, 45, 3}

	max := arr[0]
	min := arr[0]

	for _, v := range arr {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	fmt.Println("Array =", arr)
	fmt.Println("max=", max)
	fmt.Println("min=", min)
}
