package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	x := 10
	for x > 5 {
		if x%2 == 0 {
			fmt.Println(x)
		}
		x--
	}
	names := [4]string{"Lewis", "George", "Max", "Checko"}

	for i, value := range names {
		fmt.Println(i, value)
	}
}
