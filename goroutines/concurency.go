package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	waitGroup.Add(1)
	go printFirst()
	waitGroup.Wait()
	printLater()

}

func printFirst() {
	for i := 0; i < 50; i++ {
		fmt.Print(i, "_")
	}
	waitGroup.Done()
}
func printLater() {
	for i := 100; i < 130; i++ {
		fmt.Print(i, "_")
	}
}
