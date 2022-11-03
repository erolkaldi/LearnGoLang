package main

import "fmt"

func sayHi(chn chan string) {
	chn <- "Hi Github"
}

func main() {

	channel := make(chan string)
	go sayHi(channel)

	fmt.Println(<-channel)
}
