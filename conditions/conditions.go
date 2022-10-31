package main

import "fmt"

func main() {

	x, y := 5, 3

	fmt.Println(x, y)
	fmt.Printf("%T, %v\n", x == y, x == y)

	if x > y {
		fmt.Println("If condition returned ", x > y)
	}

	if x > y {
		fmt.Println("Greater")
	} else if x == y {
		fmt.Println("Equal")
	} else {
		fmt.Println("Smaller")
	}

	a, b := "A", "B"

	fmt.Println(a != b)
}
