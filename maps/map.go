package main

import "fmt"

func main() {
	maps := map[string]int{
		"Ankara":   6,
		"Antalya":  7,
		"İstanbul": 34,
		"İzmir":    35,
	}

	fmt.Println(maps["Antalya"])

	delete(maps, "İzmir")
	fmt.Println(maps)

	value, ok := maps["Konya"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("Not found")
	}

	maps["Konya"] = 42
	fmt.Println(maps)

	val, ok := maps["Konya"]
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("Not found")
	}
}
