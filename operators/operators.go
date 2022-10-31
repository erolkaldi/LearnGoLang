package main

import "fmt"

func main() {

	x, y := 5, 3
	var z = x + y
	fmt.Println(z)

	z = x - y
	fmt.Println(z)

	z = x * y
	fmt.Println(z)

	z = x / y
	fmt.Println(z)

	z = x % y
	fmt.Println(z)

	var q = 5.0 / 3.0
	fmt.Println(q)

	z = 5
	z++
	fmt.Println(z)

	z = 5
	z--
	fmt.Println(z)
}
