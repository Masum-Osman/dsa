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

func IsPalindromeTwoPointer(s string) bool {
	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		return false
	}

	sanitizedStr := reg.ReplaceAllString(s, "")
	sanitizedStr = strings.ToLower(sanitizedStr)

	start := 0
	end := len(sanitizedStr) - 1

	for start < end {
		if sanitizedStr[start] != sanitizedStr[end] {
			return false
		} else {
			start++
			end--
		}
	}

	return true
}

func IsAlphaNumeric(chr byte) bool {
	if (chr >= 48 && chr <= 57) ||
		(chr >= 65 && chr <= 90) ||
		(chr >= 97 && chr <= 122) {
		return true
	}
	return false
}

func ToLowerCase(chr byte) byte {
	if chr > 64 && chr < 91 {
		chr = chr + 32
	}

	return chr
}

func IsPalindromeEx1(s string) bool {
	for i, j := 0, len(s)-1; i < len(s) && j >= 0; {
		chrLeft := s[i]
		chrRight := s[j]

		if !IsAlphaNumeric(chrLeft) {
			i++
			continue
		}

		if !IsAlphaNumeric(chrRight) {
			j--
			continue
		}

		if ToLowerCase(chrLeft) != ToLowerCase(chrRight) {
			return false
		}

		i++
		j--
	}

	return true
}

func IsPalindromeEx2(s string) bool {
	var str strings.Builder
	for i := range s {
		if s[i] >= 65 && s[i] <= 90 {
			str.WriteByte(s[i] + 32)
		}
		if s[i] >= 97 && s[i] <= 122 {
			str.WriteByte(s[i])
		}
		if s[i] >= 48 && s[i] <= 57 {
			str.WriteByte(s[i])
		}
	}
	s = str.String()
	l, r := 0, len(s)-1
	for l <= r {
		if s[l] != s[r] {
			return false
		}
		l, r = l+1, r-1
	}
	return true
}
