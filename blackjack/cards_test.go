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
	}

	for _, test := range tests {
		t.Logf("Testing value of card %s", test.description)
		if test.got != test.want {
			t.Errorf("Incorrect card value ... got %v, want %v", test.got, test.want)
		}
	}
}
