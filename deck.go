package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create type deck that's a slice of strings

type deck []string

// Should optimize this to less than o(n^2) time
func newDeck() deck {
	cards := deck{}
	cardValues := []string{"Ace", "Two", "Three", "Four", "King"}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// Reciever
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func dealCard(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveDeckToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func readDeckFromfile(filename string) deck {
	byteslice, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return deck(strings.Split(string(byteslice), ","))
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	for i := range d {
		newPosition := rand.New(source).Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
