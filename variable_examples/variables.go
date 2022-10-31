package main

import "fmt"

func main() {

	var name string
	name = "GitHub"
	fmt.Println("Hello", name)

	var year int
	year = 2022
	fmt.Println("Year : ", year)

	var alive bool
	alive = true
	fmt.Println("Alive : ", alive)

	var fun = "have fun"
	fmt.Println(fun)

	more := "more fun"
	fmt.Println(more)

	much, more, fun := "Get", "more", "fun"
	fmt.Println(much, more, fun)

}
