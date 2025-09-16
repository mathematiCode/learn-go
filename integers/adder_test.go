package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("Add two numbers", func(t *testing.T) {
		expected := 5
		actual := Add(2, 3)

		if actual != expected {
			t.Errorf("expected %d but got %d", expected, actual)
		}
	})
}

func ExampleAdd() {
	sum := Add(2, 5)
	fmt.Println(sum)
	// Output: 7
}
