package main

import (
	"fmt"
)

func main() {
	var length, width float64

	fmt.Print("Enter the length of the rectangle: ")
	fmt.Scanln(&length)

	fmt.Print("Enter the width of the rectangle: ")
	fmt.Scanln(&width)

	area := length * width
	fmt.Printf("The area of the rectangle is: %.2f\n", area)
}
