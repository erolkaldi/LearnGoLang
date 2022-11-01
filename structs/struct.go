package main

import "fmt"

func main() {

	var kid child
	kid.name = "Joe"
	kid.age = 6
	fmt.Println(kid)

	var adam human
	adam.name = "Adam"
	adam.age = 40
	adam.gender = "Male"
	adam.isMarried = true
	adam.specs = spec{weight: 95, height: 185}
	adam.childeren = []child{
		{name: "Tom", age: 7}, {name: "Kate", age: 5},
	}
	fmt.Println(adam)

	tom := child{name: "Tom", age: 7}
	kate := child{name: "Kate", age: 5}

	eve := human{
		name:      "Eve",
		age:       35,
		gender:    "Female",
		isMarried: true,
		specs:     spec{weight: 60, height: 170},
		childeren: []child{
			tom, kate,
		},
	}
	fmt.Println(eve)
}

type child struct {
	name string
	age  int
}
type spec struct {
	weight int
	height int
}
type human struct {
	name      string
	age       int
	gender    string
	isMarried bool
	specs     spec
	childeren []child
}
