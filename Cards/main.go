package main

import "fmt"

// We can initialize a variable outside of a function, we just can't assign a value to it.
var deckSize int

// deckSize := 20 // error

func main() {
	deckSize = 50
	// var card string = "Ace of Spades"
	card := "Ace of Spades"

	card = "Five of Diamonds"
	fmt.Println(card)
	fmt.Println(deckSize)
}
