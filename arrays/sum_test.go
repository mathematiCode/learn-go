package arrays

import (
	"testing"
)

func TestSum3(t *testing.T) {
	nums := [3]int{2, 3, 4}
	got := Sum3(nums)
	want := 9
	if got != want {
		t.Errorf("got %d but wanted %d, given %v", got, want, nums)
	}
}

func TestSum(t *testing.T) {
	t.Run("sum 4 numbers", func(t *testing.T) {
		nums := []int{2, 3, 4, 5}
		got := Sum(nums)
		want := 14

		if got != want {
			t.Errorf("got %d but wanted %d, given %v", got, want, nums)
		}
	})

	t.Run("sum 2 numbers", func(t *testing.T) {
		nums := []int{2, 8}
		got := Sum(nums)
		want := 10

		if got != want {
			t.Errorf("got %d but wanted %d, given %v", got, want, nums)
		}
	})
}
