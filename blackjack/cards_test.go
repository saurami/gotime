package blackjack

import "testing"

func TestCardValue(t *testing.T) {
	tests := []struct{
		description string
		got int
		want int
	}{
		{"ace", parseCard("ace"), 11},
		{"two", parseCard("two"), 2},
		{"three", parseCard("three"), 3},
		{"four", parseCard("four"), 4},
		{"five", parseCard("five"), 5},
		{"six", parseCard("six"), 6},
		{"seven", parseCard("seven"), 7},
		{"eight", parseCard("eight"), 8},
		{"nine", parseCard("nine"), 9},
		{"king", parseCard("king"), 10},
		{"demo", parseCard("demo"), 0},
	}

	for _, test := range tests {
		t.Logf("Testing value of card %s", test.description)
		if test.got != test.want {
			t.Errorf("Incorrect card value ... got %v, want %v", test.got, test.want)
		}
	}
}
