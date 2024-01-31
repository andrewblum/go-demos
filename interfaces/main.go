package main

import "fmt"

// must list all the types of input and output for these functions
// interfaces about behavior more than data

type veryAgreeable interface{}

type animal interface {
	speak() string
}

type cat struct {
	name string
	age  int
}

type dog struct {
	name string
	age  int
}

func main() {
	dorito := cat{"dorito", 5}
	winter := dog{"winter", 7}

	// seeing what a interface value looks like
	var animalInterface animal
	fmt.Println(animalInterface)

	animalInterface = dorito
	fmt.Println(animalInterface)
	fmt.Printf("(%v, %T)\n", animalInterface, animalInterface)

	animalInterface = winter
	fmt.Println(animalInterface)
	fmt.Printf("(%v, %T)\n", animalInterface, animalInterface)

	// nil interface values
	var emptyCatPointer *cat
	animalInterface = emptyCatPointer
	fmt.Printf("(%v, %T)\n", animalInterface, animalInterface)

	// check types an interface is holding
	animalInterface = dorito
	t, ok := animalInterface.(cat)
	fmt.Println(t, ok)

	// printGettingCat(dorito)
	// printGettingDog(winter)
	// printGreeting(dorito)
	// printGreeting(winter)
	// printGreetingAnyone(winter)
	fmt.Println(dorito)
	fmt.Println(winter)

}

// notice two functions with the same name are ok
// if they are different receivers

func (cat) speak() string {
	return "meow"
}

func (dog) speak() string {
	return "woof"
}

func (d dog) String() string {
	return fmt.Sprintf("Dog named %v , and is (%d) years old", d.name, d.age)
}

// cant have two functions with the same name :)
func printGettingDog(d dog) {
	fmt.Println(d.speak())
}

func printGettingCat(c cat) {
	fmt.Println(c.speak())
}

func printGreeting(a animal) {
	fmt.Println(a.speak())
}

func printGreetingAnyone(a veryAgreeable) {
	fmt.Println(a)
}
