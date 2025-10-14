package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

//Method
func (p Person) Print() {
	fmt.Printf("Name:%s,Age:%d\n", p.Name, p.Age)
}

func main() {
	// create instance

	p1 := Person{Name: "Tibin", Age: 25}

	// calling method

	p1.Print()
}
