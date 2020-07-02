package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected 16, got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected Four of Clubs, got %v", d[len(d)-1])
	}
}

func TestAndBlabla(t *testing.T) {
	fn := "__decktesting"
	os.Remove(fn)
	deck := newDeck()
	deck.saveToFile(fn)

	loaded := newDeckFromFile(fn)
	if len(loaded) != 16 {
		t.Errorf("expected 16, got %v", len(loaded))
	}
	os.Remove(fn)
}
