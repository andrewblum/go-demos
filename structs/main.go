package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	andrew := person{"Andrew", "Blum", contactInfo{"a", 1}}
	andrew2 := person{firstName: "Andrew", lastName: "Blum"}
	var andrew3 person
	andrew3.firstName = "Andrew"
	andrew3.lastName = "Blum"
	fmt.Println(andrew)
	fmt.Println(andrew2)
	fmt.Println(andrew3)
	andrew3.print()
	andrew3pointer := &andrew3
	andrew3pointer.updateName("andy")
	andrew3.print()
	andrew3.updateNamePointerShortCut("short andy")
	andrew3.print()

	// this will error!
	// fmt.Println(andrew3.donut)

	// MAPS

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	//colors2 := make(map[string]string)
	// access with bracket notation
	// delete(map, key)
	colors["blue"] = "#0000FF"
	for key, val := range colors {
		fmt.Println(key, val)
	}

}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p *person) updateNamePointerShortCut(newFirstName string) {
	(*p).firstName = newFirstName
}
