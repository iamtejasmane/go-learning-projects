package main

func main() {
	//cards := newDeckFromFile("my_cards")
	cards := newDeck()

	cards.shuffle()
	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
