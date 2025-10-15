package main

import (
	"errors"
	"fmt"
)

// Struct: Student
type Student struct {
	Name string
	Age  int
}

// Slice to store students
var students []Student

// AddStudent - Adds a new student to the slice
func AddStudent(name string, age int) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	if age <= 0 {
		return errors.New("age must be positive")
	}
	students = append(students, Student{Name: name, Age: age})
	fmt.Println(" Student added successfully!")
	return nil
}

// ListStudents - Displays all students
func ListStudents() {
	if len(students) == 0 {
		fmt.Println("  No students found.")
		return
	}
	fmt.Println("\n List of Students:")
	for i, s := range students {
		fmt.Printf("[%d] %s (%d years old)\n", i, s.Name, s.Age)
	}
}

// UpdateStudent - Updates student details using index
func UpdateStudent(index int, newName string, newAge int) error {
	if index < 0 || index >= len(students) {
		return errors.New("invalid student index")
	}
	if newName == "" {
		return errors.New("name cannot be empty")
	}
	if newAge <= 0 {
		return errors.New("age must be positive")
	}
	students[index].Name = newName
	students[index].Age = newAge
	fmt.Println(" Student updated successfully!")
	return nil
}

// DeleteStudent - Removes a student by index
func DeleteStudent(index int) error {
	if index < 0 || index >= len(students) {
		return errors.New("invalid student index")
	}
	students = append(students[:index], students[index+1:]...)
	fmt.Println(" Student deleted successfully!")
	return nil
}

func main() {
	for {
		fmt.Println("   Student Management System   ")
		fmt.Println("1. Add Student")
		fmt.Println("2. List Students")
		fmt.Println("3. Update Student")
		fmt.Println("4. Delete Student")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println(" Invalid input. Please enter a number.")
			fmt.Scanln() // clear buffer
			continue
		}

		switch choice {
		case 1:
			var name string
			var age int
			fmt.Print("Enter student name: ")
			fmt.Scan(&name)
			fmt.Print("Enter student age: ")
			fmt.Scan(&age)

			err := AddStudent(name, age)
			if err != nil {
				fmt.Println(" Error:", err)
			}

		case 2:
			ListStudents()

		case 3:
			if len(students) == 0 {
				fmt.Println(" No students to update.")
				continue
			}
			ListStudents()

			var index int
			fmt.Print("Enter student index to update: ")
			fmt.Scan(&index)

			var newName string
			var newAge int
			fmt.Print("Enter new name: ")
			fmt.Scan(&newName)
			fmt.Print("Enter new age: ")
			fmt.Scan(&newAge)

			err := UpdateStudent(index, newName, newAge)
			if err != nil {
				fmt.Println("❌ Error:", err)
			}

		case 4:
			if len(students) == 0 {
				fmt.Println("⚠️ No students to delete.")
				continue
			}
			ListStudents()

			var index int
			fmt.Print("Enter student index to delete: ")
			fmt.Scan(&index)

			err := DeleteStudent(index)
			if err != nil {
				fmt.Println(" Error:", err)
			}

		case 5:
			fmt.Println(" Exiting program... Goodbye!")
			return

		default:
			fmt.Println(" Invalid choice. Please select between 1–5.")
		}
	}
}
