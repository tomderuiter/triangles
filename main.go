package main

import (
	"fmt"

	triangles "github.com/tomderuiter/triangles/triangles"
)

func main() {
	fmt.Println("I will help you classify the right triangle by side. Please enter the length of all three sides.")
	var A, B, C int
	for {

		fmt.Print("Enter the first number: ")
		fmt.Scan(&A)

		fmt.Print("Enter the second number: ")
		fmt.Scan(&B)

		fmt.Print("Enter the third number: ")
		fmt.Scan(&C)

		triangles.ClassifyTriangle(A, B, C)

		var answer string

		fmt.Println("Do you want to classify another triangle? (yes/no)")
		fmt.Scan(&answer)

		if answer == "yes" {

		} else {
			break

		}

	}
}
