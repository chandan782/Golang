//composition means putting together
//In golang composition can be achieved by embedding one struct type into another, which is called as embedding.

package main

import "fmt"

type Vehicle struct {
	Name string
	Type string
}

type Car struct {
	Vehicle
	MaxSpeed float32
	FuelType string
}

func main() {
	c := Car{}
	c.Name = "Ferrari"
	c.Type = "Convertible"
	c.MaxSpeed = 250
	c.FuelType = "Diesel"

	fmt.Println(c)

}


//This indicates that car has vehicle which means its a composition
