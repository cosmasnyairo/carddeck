package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 20 {
		t.Errorf("Expected deck length of 20 but got: %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of `Ace of spades` but got: %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card of `King of Clubs` but got: %v", d[len(d)-1])
	}

}

func TestSaveDeckToFileReadDeckFromfile(t *testing.T) {

	d := newDeck()
	d.saveDeckToFile("deck_test.txt")

	loadedDeck := readDeckFromfile("deck_test.txt")

	if len(loadedDeck) != 20 {
		t.Errorf("Expected deck length of 20 but got: %v", len(loadedDeck))
	}

	if loadedDeck[0] != "Ace of Spades" {
		t.Errorf("Expected first card of `Ace of spades` but got: %v", loadedDeck[0])
	}

	if loadedDeck[len(loadedDeck)-1] != "King of Clubs" {
		t.Errorf("Expected last card of `King of Clubs` but got: %v", d[len(loadedDeck)-1])
	}
	os.Remove("deck_test.txt")
}
