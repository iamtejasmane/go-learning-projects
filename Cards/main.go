package main

// func main() {
// 	card := newCard()

// 	fmt.Println(card)
// }

func newCard() string {
	return "Five of Diamonds"
}

// Slice
func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()
}

// type book string

// func (b book) printBook() {
// 	fmt.Println(b)
// }

// func main() {
// 	var b book = "Harry Potter"
// 	b.printBook()
// }
