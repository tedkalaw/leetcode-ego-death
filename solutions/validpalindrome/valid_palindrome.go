package validpalindrome

import (
	"strings"
	"unicode"
)

// Valid Palindrome
//
// Given a string s, return true if it is a palindrome, otherwise return false.
//
// A palindrome is a string that reads the same forward and backward.
// It is also case-insensitive and ignores all non-alphanumeric characters.
//
// Example 1:
// Input: s = "Was it a car or a cat I saw?"
// Output: true
// Explanation: After considering only alphanumerical characters we have "wasitacaroracatisaw", which is a palindrome.
//
// Example 2:
// Input: s = "tab a cat"
// Output: false
// Explanation: "tabacat" is not a palindrome.
//
// Constraints:
// 1 <= s.length <= 1000
// s is made up of only printable ASCII characters.

func TwoPointerSolution(s string) bool {
	if len(s) == 0 {
		return true
	}

	// Go's standard library is honestly pretty sweet
	lowercaseStr := strings.ToLower(s)

	// Pre-allocating a slice with a capacity of lowercaseStr means
	// that we know that when we use append, we won't be causing an
	// extra allocation to happen
	// Prolly not a big deal anyway since append is presumably doubling
	// size each time, but still
	var alphanumeric []rune = make([]rune, 0, len(lowercaseStr))

	for _, r := range lowercaseStr {
		// The problem actually specifics that it's ascii instead of unicode, but
		// it's pretty cool to use the unicode package
		if !unicode.IsDigit(r) && !unicode.IsLetter(r) {
			continue
		}

		alphanumeric = append(alphanumeric, r)
	}

	if len(alphanumeric) == 0 {
		return true
	}

	left, right := 0, len(alphanumeric)-1

	for left < right {
		if alphanumeric[left] != alphanumeric[right] {
			return false
		}

		left++
		right--
	}

	return true
}
