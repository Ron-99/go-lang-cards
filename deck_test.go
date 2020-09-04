package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	const LENDECK = 16

	if len(d) != LENDECK {
		t.Errorf("Expected deck length of %d, but got %d", LENDECK, len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}
