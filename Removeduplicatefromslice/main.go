package main

import (
	"fmt"
)

func main() {
	nums := []int{10, 20, 30, 10, 20, 40, 50, 40}

	uniqueMap := make(map[int]bool) //create a map
	result := []int{}               //create a empty slice for unique numbers

	for _, v := range nums {
		if !uniqueMap[v] {
			uniqueMap[v] = true
			result = append(result, v)
		}
	}

	fmt.Println(nums)
	fmt.Println(result)
}
