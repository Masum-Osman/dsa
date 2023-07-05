package twoPointer

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	type TestCases struct {
		Input  string
		Output bool
	}

	cases := []TestCases{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
	}

	for i, tc := range cases {
		got := IsPalindromeTwoPointer(tc.Input)

		if got != tc.Output {
			t.Error("For test Case: ", i+1, "Wanted: ", tc.Output, " || Got: ", got)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	type TestCases struct {
		Input  string
		Output bool
	}

	cases := []TestCases{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
	}

	for i, tc := range cases {
		got := IsPalindromeEx2(tc.Input)

		if got != tc.Output {
			b.Error("For test Case: ", i+1, "Wanted: ", tc.Output, " || Got: ", got)
		}
	}
}
