package main

import (
	"fmt"
	"time"
)

func main() {

	var (
		name   string = "Hello String"
		year   int    = 2022
		answer bool   = false
	)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", year)
	fmt.Printf("%T\n", answer)

	var age byte = 40
	fmt.Printf("%T\n", age)

	var height float32 = 190.20
	fmt.Printf("%T\n", height)

	var today = time.Now()
	today = time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, time.UTC)
	fmt.Printf("%T\n", today)
	fmt.Println(today)

}
