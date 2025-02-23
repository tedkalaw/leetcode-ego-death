package tests

import (
	"testing"

	"github.com/tedkalaw/leetcode-ego-death/solutions"
)

func TestValidAnagram(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "Basic anagram",
			s:    "racecar",
			t:    "carrace",
			want: true,
		},
		{
			name: "Different words",
			s:    "jar",
			t:    "jam",
			want: false,
		},
		{
			name: "Empty strings",
			s:    "",
			t:    "",
			want: true,
		},
		{
			name: "Single character same",
			s:    "a",
			t:    "a",
			want: true,
		},
		{
			name: "Single character different",
			s:    "a",
			t:    "b",
			want: false,
		},
		{
			name: "Different lengths",
			s:    "rat",
			t:    "car",
			want: false,
		},
		{
			name: "Complex anagram",
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{
			name: "Same letters different count",
			s:    "aa",
			t:    "a",
			want: false,
		},
		{
			name: "Long anagram",
			s:    "debitcard",
			t:    "badcredit",
			want: true,
		},
		{
			name: "Similar but not anagram",
			s:    "cat",
			t:    "rat",
			want: false,
		},
		{
			name: "Repeated letters",
			s:    "hello",
			t:    "llohe",
			want: true,
		},
		{
			name: "All same letter",
			s:    "aaaa",
			t:    "aaaa",
			want: true,
		},
		{
			name: "Different count of same letter",
			s:    "aaa",
			t:    "aaaa",
			want: false,
		},
		{
			name: "Complex with spaces (not anagram)",
			s:    "hello world",
			t:    "world hello",
			want: true,
		},
		{
			name: "Multiple repeated letters",
			s:    "mississippi",
			t:    "sipssisipmi",
			want: true,
		},
		{
			name: "Almost anagram",
			s:    "paper",
			t:    "piper",
			want: false,
		},
		{
			name: "Different case (not anagram by problem definition)",
			s:    "Hello",
			t:    "hello",
			want: false,
		},
		{
			name: "Long identical strings",
			s:    "abcdefghijklmnopqrstuvwxyz",
			t:    "abcdefghijklmnopqrstuvwxyz",
			want: true,
		},
		{
			name: "Palindrome anagram",
			s:    "racecar",
			t:    "racecar",
			want: true,
		},
		{
			name: "Different characters same length",
			s:    "abcde",
			t:    "vwxyz",
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := solutions.ValidAnagram(tc.s, tc.t)
			if got != tc.want {
				t.Errorf("ValidAnagram(%v, %v) = %v; want %v", tc.s, tc.t, got, tc.want)
			}
		})
	}
}
