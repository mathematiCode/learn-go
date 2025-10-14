package other

import "testing"

func TestContrast(t *testing.T) {
	got, err := MaxContrast("#A1A3F7")
	want := "#000000"

	if err != nil {
		t.Errorf("There was an error: %s", err)
	}

	if got != want {
		t.Errorf("got %s but wanted %s", got, want)
	}
}
