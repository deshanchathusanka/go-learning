package main

import "testing"

func TestBuildDeck(t *testing.T) {
	deck := buildDeck()
	if len(deck) != 16 {
		t.Errorf("Expected 16, But received %v", len(deck))
	}
}
