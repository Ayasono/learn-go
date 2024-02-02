package main

func main() {
	cards := newDeck()

	cards.shuffle()

	err := cards.saveToFile("cards_cache")
	if err != nil {
		return
	}
	cards.print()
}
