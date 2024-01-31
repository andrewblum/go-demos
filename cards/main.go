package main

import "fmt"

// reciver functions can be declared on any type
type donut int

func (d donut) eat() {
	fmt.Println("yummy")
}

func main() {

	var delicious_circle donut = 8
	delicious_circle2 := donut(7)
	delicious_circle.eat()
	delicious_circle2.eat()

	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)
	fmt.Println("***")
	hand.print()
	fmt.Println("***")
	remainingDeck.print()
	fmt.Println("***")
	fmt.Println(cards)
	fmt.Println(cards.toString())

	cards.saveToFile("my_cards")
	d := newDeckFromFile("my_cards")
	fmt.Println(d)

}
