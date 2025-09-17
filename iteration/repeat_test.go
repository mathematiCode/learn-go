package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {

	got := Repeat("a", 5)
	want := "aaaaa"
	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}
