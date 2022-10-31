package main

import "fmt"

func main() {

	sayHello("Github")
	a, b := 5, 3

	var c = sum(a, b)

	fmt.Printf("%d and %d summary %d\n", a, b, c)
	fmt.Printf("%d and %d summary %d\n", b, c, sum(b, c))
	fmt.Println(checkNumber(a))
	fmt.Println(checkNumber(b))
	fmt.Println(checkNumber(sum(b, c)))

	m, n := division(a, b)
	fmt.Printf("%d and %d division result is %d and rest is %d", a, b, m, n)

}
func sum(x, y int) int {
	return x + y
}
func sayHello(name string) {
	fmt.Println("Hello", name)
}

func checkNumber(x int) string {

	if x > 5 {
		return "Greater than 5"
	}
	if x == 5 {
		return "Equal to 5"
	}
	return "Smaller than 5"
}

func division(value, divider int) (result, rest int) {
	result = value / divider
	rest = value % divider
	return result, rest
}
