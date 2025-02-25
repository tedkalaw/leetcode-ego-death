package tests

import (
	"testing"

	"github.com/tedkalaw/leetcode-ego-death/solutions/validpalindrome"
)

func TestValidPalindrome(t *testing.T) {
	solutions := map[string]func(string) bool{
		"twopointer": validpalindrome.TwoPointerSolution,
	}

	testCases := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "Example1",
			input: "Was it a car or a cat I saw?",
			want:  true,
		},
		{
			name:  "Example2",
			input: "tab a cat",
			want:  false,
		},
		{
			name:  "Empty string",
			input: "",
			want:  true,
		},
		{
			name:  "Single character",
			input: "a",
			want:  true,
		},
		{
			name:  "Simple palindrome",
			input: "racecar",
			want:  true,
		},
		{
			name:  "Simple non-palindrome",
			input: "hello",
			want:  false,
		},
		{
			name:  "Palindrome with spaces",
			input: "A man a plan a canal Panama",
			want:  true,
		},
		{
			name:  "Palindrome with mixed case",
			input: "RaceCar",
			want:  true,
		},
		{
			name:  "Palindrome with punctuation",
			input: "A man, a plan, a canal: Panama",
			want:  true,
		},
		{
			name:  "Palindrome with numbers",
			input: "12321",
			want:  true,
		},
		{
			name:  "Mixed alphanumeric palindrome",
			input: "A1b2c3c2b1a",
			want:  true,
		},
		{
			name:  "Mixed alphanumeric non-palindrome",
			input: "A1b2c3d4e5",
			want:  false,
		},
		{
			name:  "Palindrome with special characters",
			input: "No 'x' in Nixon",
			want:  true,
		},
		{
			name:  "Long palindrome",
			input: "Doc, note: I dissent. A fast never prevents a fatness. I diet on cod.",
			want:  true,
		},
		{
			name:  "Almost palindrome",
			input: "racecars",
			want:  false,
		},
		{
			name:  "Palindrome with unicode",
			input: "A Santa at NASA",
			want:  true,
		},
		{
			name:  "Only special characters (palindrome)",
			input: "!@#$@!",
			want:  true,
		},
		{
			name:  "Palindrome with mixed spaces",
			input: "never odd or even",
			want:  true,
		},
		{
			name:  "Palindrome with digits and punctuation",
			input: "1 eye for of 1 eye.",
			want:  false,
		},
		{
			name:  "Palindrome with repeated characters",
			input: "AAAAA",
			want:  true,
		},
		{
			name:  "Palindrome with repeated characters and spaces",
			input: "A A A A A",
			want:  true,
		},
		{
			name:  "Palindrome with repeated words",
			input: "bob bob",
			want:  true,
		},
		{
			name:  "Non-palindrome with repeated words",
			input: "bob alice bob",
			want:  false,
		},
		{
			name:  "Palindrome with hyphen",
			input: "A man, a plan, a canal - Panama",
			want:  true,
		},
		{
			name:  "Palindrome with underscore",
			input: "race_car",
			want:  true,
		},
		{
			name:  "Palindrome with exclamation",
			input: "Madam, I'm Adam!",
			want:  true,
		},
		{
			name:  "Palindrome with question mark",
			input: "Was it a rat I saw?",
			want:  true,
		},
		{
			name:  "Palindrome with quotes",
			input: "\"Able was I ere I saw Elba\"",
			want:  true,
		},
		{
			name:  "Palindrome with parentheses",
			input: "(Able was I ere I saw Elba)",
			want:  true,
		},
		{
			name:  "All non-alphanumeric",
			input: ".,;:!@#$%^&*()",
			want:  true,
		},
	}

	for name, solution := range solutions {
		t.Run(name, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					got := solution(tc.input)
					if got != tc.want {
						t.Errorf("ValidPalindrome(%q) = %v; want %v", tc.input, got, tc.want)
					}
				})
			}
		})
	}
}
