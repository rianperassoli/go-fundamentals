package main

func main() {
	// var card string = "Ace of spades"
	cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")

	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "Five of diamonds"
}
