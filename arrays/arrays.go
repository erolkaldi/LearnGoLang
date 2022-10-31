package main

import "fmt"

func main() {
	words := [2]string{"Hello", "Github"}
	fmt.Println(words)
	fmt.Println(words[0])
	fmt.Println(words[1])
	fmt.Println("Count ", len(words))

	numbers := [...]int{1, 2, 3, 4, 5}

	fmt.Println(numbers)
}
