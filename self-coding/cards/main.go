package main

func main() {
	// cards := newDeckFromFile("cards")
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
