package techpalace

import "testing"

func TestWelcomeMessage(t *testing.T) {
	got := welcomeMessage("Judy")
	want := "Welcome to the Tech Palace, JUDY"

	if got != want {
		t.Errorf("Error with welcome message ... got %v; wanted %v", got, want)
	}
}

func TestBorderMessage(t *testing.T) {
	got := addBorder("Welcome", 10)
	want :=`**********
Welcome
**********`

	if got != want {
		t.Errorf("Error with display panel ... got \n%v\n; wanted \n%v\n", got, want)
	}
}

func TestCleanupMessage(t *testing.T) {
	message := `
**************************
*    BUY NOW, SAVE 10%   *
**************************`
	got := cleanupMessage(message)
	want := "BUY NOW, SAVE 10%"

	if got != want {
		t.Errorf("Error cleaning up message ... got \n%v\n; wanted \n%v\n", got, want)
	}
}
