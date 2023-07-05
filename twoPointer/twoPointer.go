package twoPointer

import (
	"regexp"
	"strings"
)

/*
LC : 125
*/

func ReverseString(s string) (reversed string) {
	for _, ch := range s {
		reversed = string(ch) + reversed
	}

	return
}

func IsPalindrome(s string) bool {
	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		return false
	}

	sanitizedStr := reg.ReplaceAllString(s, "")
	sanitizedStr = strings.ToLower(sanitizedStr)

	reversedStr := ReverseString(sanitizedStr)

	return reversedStr == sanitizedStr
}
