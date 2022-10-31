package main

import "fmt"

func main() {
	var baseSlice = []int{1, 2, 3}
	fmt.Println(baseSlice)
	var slice []int
	slice = make([]int, 3)
	fmt.Println(slice)
	slice = baseSlice
	fmt.Println(slice)
	slice[0] = 4
	fmt.Println(slice)
	fmt.Println(baseSlice)
	fmt.Println("Base slice also changed")
	fmt.Println(slice[:])
	fmt.Println(slice[2:])
	fmt.Println(slice[1:3])

}
