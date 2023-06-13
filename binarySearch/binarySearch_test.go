package binarySearch

import "testing"

func TestSearch(t *testing.T) {
	type testCases struct {
		nums   []int
		target int
		result int
	}

	cases := []testCases{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{[]int{-1, 0, 3, 5, 9, 12}, -1, 0},
		{[]int{-1, 0, 3, 5, 9, 12}, 12, 5},
	}

	for _, tc := range cases {
		got := Search(tc.nums, tc.target)

		t.Log(got)
		if tc.result != got {
			t.Error("Wanted: ", tc.result, "Got: ", got)
		}
	}
}

func TestSearchInsert(t *testing.T) {
	type testCases struct {
		nums   []int
		target int
		result int
	}

	cases := []testCases{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 2, 3, 4, 5, 10}, 2, 1},
		{[]int{1, 2, 3, 4, 5, 10, 11}, 2, 1},
		{[]int{1, 2, 3, 4, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 0, 0},
		{[]int{1, 3}, 2, 1},
	}

	for _, tc := range cases {
		got := SearchInsertUpdated(tc.nums, tc.target)

		t.Log(got)
		if tc.result != got {
			t.Error("Wanted: ", tc.result, "Got: ", got)
		}
	}
}
