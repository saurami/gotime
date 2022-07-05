package greet

import "testing"

func TestHello(t *testing.T) {
	got := message()
	want := "Hello, World!"

	if got != want {
		t.Errorf("Got %q; Want %q", got, want)
	}
}
