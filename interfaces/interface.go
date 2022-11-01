package main

import "fmt"

type car struct {
	name      string
	km        int
	fuelPerKm int
}

func (c car) efficency() int {
	return c.km / c.fuelPerKm
}

type truck struct {
	name      string
	km        int
	fuelPerKm int
	load      int
}

func (t truck) efficency() int {
	return (t.km * t.load) / t.fuelPerKm
}

type vehicle interface {
	efficency() int
}

func efficency(v vehicle) int {
	return v.efficency()
}
func main() {
	volvo := car{
		name:      "VolVo",
		km:        500,
		fuelPerKm: 2,
	}
	dodge := truck{
		name:      "Dodge",
		km:        700,
		fuelPerKm: 3,
		load:      40,
	}
	fmt.Println(efficency(volvo))
	fmt.Println(efficency(dodge))
}
