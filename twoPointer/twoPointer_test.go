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
		got := IsPalindrome(tc.Input)

		if got != tc.Output {
			t.Error("For test Case: ", i+1, "Wanted: ", tc.Output, " || Got: ", got)
		}
	}
}
