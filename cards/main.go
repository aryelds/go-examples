package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	//cards := newDeckFromFile("my_cardwwws")
	//cards.print()

	//cards := newDeck()
	//cards.saveToFile("my_cards")

}
