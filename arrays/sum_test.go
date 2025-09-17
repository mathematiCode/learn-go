package arrays

import (
	"reflect"
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

	checksums := func(t *testing.T, got, want int) {
		if got != want {
			t.Errorf("got %+v but wanted %+v", got, want)
		}
	}
	t.Run("sum 4 numbers", func(t *testing.T) {
		nums := []int{2, 3, 4, 5}
		got := Sum(nums)
		want := 14
		checksums(t, got, want)
	})

	t.Run("sum 2 numbers", func(t *testing.T) {
		nums := []int{2, 8}
		got := Sum(nums)
		want := 10
		checksums(t, got, want)
	})
}

func TestSumAll(t *testing.T) {
	checksums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v but wanted %#v", got, want)
		}
	}

	nums1 := []int{3, 4}
	nums2 := []int{1, 2, 5, 10}

	got := SumAll(nums1, nums2)
	want := []int{7, 18}
	checksums(t, got, want)
}
