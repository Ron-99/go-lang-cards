package main

func main() {
	cards := newDeck()
	
	hand, _ := deal(cards, 5)

	hand.shuffle()
	hand.print()
}
