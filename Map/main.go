package main

import "fmt"

func main() {
	// initialize a map with students name and grades

	students := map[string]int{
		"Alice": 85,
		"Bob":   90,
		"Carol": 78,
	}
	fmt.Println("initial list", students)

	//update a student's grade
	students["Alice"] = 95

	//add a new studen
	students["Tibin"] = 70

	//delete a student
	delete(students, "Carol")

	//iterate

	for name, grade := range students {
		fmt.Printf("%s-%d\n", name, grade)
	}
	//check if a student is exists

	name := "Haheem"
	if grade, ok := students[name]; ok {
		fmt.Printf("\n%s's grade:%d\n", name, grade)
	} else {
		fmt.Printf("\n%s not found in record\n", name)
	}

}
