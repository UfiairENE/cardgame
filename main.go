package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()

	before, after := deal(cards, 4)
	fmt.Println("dealed cards : ", before, " ", after)

	after.shuffle()
	after.saveToFile("card.txt")

	after.print()

	deckFromFile := newDeckFromFile("card.txt")
	deckFromFile.print()
}
