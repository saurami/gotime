package partyrobot

import "testing"

func TestWelcome(t *testing.T) {
	got := welcome("Christiane")
	want := "Welcome to my party, Christiane!"

	if got != want {
		t.Errorf("Incorrect welcome message ... got %v; want %v", got, want)
	}
}

func TestHappyBirthday(t *testing.T) {
	got := happyBirthday("Frank", 58)
	want := "Happy birthday Frank! You are now 58 years old!"

	if got != want {
		t.Errorf("Incorrect name or age ... got %v; want %v", got, want)
	}
}

func TestAssignTable(t *testing.T) {
	var name, neighbor, direction string = "Christiane", "Frank", "on the left"
	table, distance := 27, 23.7834298
	got := assignTable(name, table, neighbor, direction, distance)
	want := `Welcome to my party, Christiane!
You have been assigned to table 027. Your table is on the left, exactly 23.8 meters from here.
You will be sitting next to Frank.
`

	if got != want {
		t.Errorf("Incorrect information ... got \n%v\n; wanted \n%v\n", got, want)
	}
}
