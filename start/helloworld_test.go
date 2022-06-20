package start

import "testing"

func TestHello(t *testing.T) {
	got := hello()
	want := "Hello, World!"

	if got != want {
		t.Errorf("Got %q; Want %q", got, want)
	}
}
