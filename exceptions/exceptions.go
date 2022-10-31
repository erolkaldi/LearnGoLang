package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	a, b := 10, 7
	result, _ := getHalf(a)
	fmt.Println(result)
	result, err := getHalf(b)

	if err != nil {
		fmt.Println(err.Error())
	}
	openFile()
}
func getHalf(x int) (int, error) {
	if x%2 != 0 {
		return 0, errors.New("Even numbers needed")
	}
	y := x / 2
	return y, nil
}

func openFile() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	fmt.Println(file.Name())
}
