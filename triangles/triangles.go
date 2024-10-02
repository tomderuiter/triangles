package triangles

import "fmt"

func ClassifyTriangle(a int, b int, c int) {
	if a <= 0 || b <= 0 || c <= 0 {
		fmt.Println("This is not a triangle. Please try again :(")

	} else if a == b && b == c {
		fmt.Println("This is an equilateral triangle")

	} else if a == b || b == c || c == a {
		fmt.Println("This is an isosceles triangle")

	} else {
		fmt.Println("This is a scalene triangle")
	}
}
