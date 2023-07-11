package array

import (
	"fmt"
	"reflect"
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
		{[]int{1, 2, 3, 1}, true},
		{[]int{0, 4, 5, 0, 3, 6}, true},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, tc := range cases {
		got := ContainsDuplicateUsingMap(tc.input)

		if tc.result != got {
			t.Error("Wanted: ", tc.result, " || Got: ", got)
		}
	}
}

func TestReverseString(t *testing.T) {
	got := ReverseString("abcdef")
	fmt.Println(got)
}

func TestValidAnagram(t *testing.T) {
	type testCases struct {
		str1   string
		str2   string
		output bool
	}

	cases := []testCases{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"anaal", "naala", true},
		{"aacc", "ccac", false},
		{"feem", "duck", false},
	}

	for _, tc := range cases {
		got := ValidAnagram(tc.str1, tc.str2)

		// fmt.Println(got)
		if tc.output != got {
			t.Error("Wanted: ", tc.output, " || Got: ", got)
		}
	}
}

func TestTwoSum(t *testing.T) {
	type testCases struct {
		nums   []int
		target int
		output []int
	}

	cases := []testCases{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, tc := range cases {
		got := TwoSum(tc.nums, tc.target)

		fmt.Println(got)
		if !reflect.DeepEqual(got, tc.output) {
			t.Error("Wanted: ", tc.output, " || Got: ", got)
		}
	}
}

func TestGroupAnagrams(t *testing.T) {
	type testCases struct {
		Input  []string
		Output [][]string
	}

	cases := []testCases{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		{[]string{""}, [][]string{{""}}},
		{[]string{"a"}, [][]string{{"a"}}},
	}

	for i, tc := range cases {
		got := GroupAnagrams(tc.Input)

		if !reflect.DeepEqual(got, tc.Output) {
			t.Error("For test Case: ", i+1, "Wanted: ", tc.Output, " || Got: ", got)
		}
	}
}

func TestTopKFrequent(t *testing.T) {
	type testCases struct {
		Input  []int
		K      int
		Output []int
	}

	cases := []testCases{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 1, 1, 2, 2, 3, 3, 3, 3, 3}, 2, []int{3, 1}},
	}

	for i, tc := range cases {
		got := TopKFrequent(tc.Input, tc.K)

		if !reflect.DeepEqual(got, tc.Output) {
			t.Error("For test Case: ", i+1, "Wanted: ", tc.Output, " || Got: ", got)
		}
	}
}
