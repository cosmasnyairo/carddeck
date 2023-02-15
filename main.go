package main

func main() {

	// var card string  = "Ace of spades"

	// var card string
	// card = "Ace of spades"

	// cardarray := [4]string{newCard(), "Six of spades", "Six of spades"}
	// cardslice := []string{newCard(), "Six of spades"}

	// hand, remainingcards := dealCard(cardslice, 4)
	// hand.print()
	// remainingcards.print()
	// fmt.Println(cardslice.toString())
	cardslice := newDeck()

	cardslice.saveDeckToFile("deck.txt")
	readdeck := readDeckFromfile("deck.txt")
	readdeck.shuffle()
	readdeck.print()
}
