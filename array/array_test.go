package array

import (
	"testing"
)

func TestSum(t *testing.T) {
	/*
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	*/

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestContainsDuplicates(t *testing.T) {
	type testCases struct {
		input  []int
		result bool
	}

	cases := []testCases{
		{[]int{7, 5, -2, -4, 0}, false},
		// {[]int{1, 2, 3, 1}, true},
		// {[]int{0, 4, 5, 0, 3, 6}, true},
		// {[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, tc := range cases {
		got := ContainsDuplicateUsingMap(tc.input)

		if tc.result != got {
			t.Error("Wanted: ", tc.result, " || Got: ", got)
		}
	}
}
