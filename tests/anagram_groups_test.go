package tests

import (
	"reflect"
	"sort"
	"testing"

	"github.com/tedkalaw/leetcode-ego-death/solutions/anagramgroups"
)

// Helper function to normalize the output for comparison
func normalizeGroups(groups [][]string) [][]string {
	// Sort strings within each group
	for _, group := range groups {
		sort.Strings(group)
	}

	// Sort groups by first string in each group
	sort.Slice(groups, func(i, j int) bool {
		return groups[i][0] < groups[j][0]
	})

	return groups
}

func TestAnagramGroups(t *testing.T) {
	solutions := map[string]func([]string) [][]string{
		"sorting": anagramgroups.SortingSolution,
	}

	testCases := []struct {
		name  string
		input []string
		want  [][]string
	}{
		{
			name:  "Example from prompt",
			input: []string{"act", "pots", "tops", "cat", "stop", "hat"},
			want:  [][]string{{"hat"}, {"act", "cat"}, {"pots", "stop", "tops"}},
		},
		{
			name:  "Single string",
			input: []string{"x"},
			want:  [][]string{{"x"}},
		},
		{
			name:  "Empty string",
			input: []string{""},
			want:  [][]string{{""}},
		},
		{
			name:  "Multiple empty strings",
			input: []string{"", "", ""},
			want:  [][]string{{"", "", ""}},
		},
		{
			name:  "No anagrams",
			input: []string{"cat", "dog", "pig"},
			want:  [][]string{{"cat"}, {"dog"}, {"pig"}},
		},
		{
			name:  "All same anagrams",
			input: []string{"eat", "ate", "tea", "eta"},
			want:  [][]string{{"ate", "eat", "eta", "tea"}},
		},
		{
			name:  "Single character strings",
			input: []string{"a", "b", "c", "a"},
			want:  [][]string{{"a", "a"}, {"b"}, {"c"}},
		},
		{
			name:  "Mixed lengths",
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want:  [][]string{{"ate", "eat", "tea"}, {"bat"}, {"nat", "tan"}},
		},
		{
			name:  "Repeated words",
			input: []string{"eat", "eat", "eat", "tea"},
			want:  [][]string{{"eat", "eat", "eat", "tea"}},
		},
		{
			name:  "Long words",
			input: []string{"debitcard", "badcredit"},
			want:  [][]string{{"badcredit", "debitcard"}},
		},
		{
			name:  "Similar but not anagrams",
			input: []string{"cat", "cats", "acts", "tacs"},
			want:  [][]string{{"acts", "cats", "tacs"}, {"cat"}},
		},
		{
			name:  "Multiple groups of pairs",
			input: []string{"ab", "ba", "cd", "dc", "ef", "fe"},
			want:  [][]string{{"ab", "ba"}, {"cd", "dc"}, {"ef", "fe"}},
		},
		{
			name:  "Single letter differences",
			input: []string{"abc", "bac", "cba", "abd"},
			want:  [][]string{{"abc", "bac", "cba"}, {"abd"}},
		},
		{
			name:  "Palindromes and non-palindromes",
			input: []string{"racecar", "carrace", "radar", "hello"},
			want:  [][]string{{"carrace", "racecar"}, {"hello"}, {"radar"}},
		},
		{
			name:  "All palindromes",
			input: []string{"mom", "dad", "wow"},
			want:  [][]string{{"dad"}, {"mom"}, {"wow"}},
		},
		{
			name:  "Repeated letters",
			input: []string{"aaa", "aaa", "bbb"},
			want:  [][]string{{"aaa", "aaa"}, {"bbb"}},
		},
		{
			name:  "Complex anagrams",
			input: []string{"algorithm", "logarithm", "logarithm"},
			want:  [][]string{{"algorithm", "logarithm", "logarithm"}},
		},
		{
			name:  "Mixed complexity",
			input: []string{"listen", "silent", "a", "a", "inlets"},
			want:  [][]string{{"a", "a"}, {"inlets", "listen", "silent"}},
		},
		{
			name:  "Sequential strings",
			input: []string{"abc", "bcd", "cde", "def"},
			want:  [][]string{{"abc"}, {"bcd"}, {"cde"}, {"def"}},
		},
		{
			name:  "Multiple single letters",
			input: []string{"a", "b", "c", "a", "b", "c"},
			want:  [][]string{{"a", "a"}, {"b", "b"}, {"c", "c"}},
		},
		{
			name:  "Mixed case (should be distinct)",
			input: []string{"eat", "Eat", "EAT"},
			want:  [][]string{{"EAT"}, {"Eat"}, {"eat"}},
		},
		{
			name:  "Very long strings",
			input: []string{"abcdefghijklmnop", "ponmlkjihgfedcba"},
			want:  [][]string{{"abcdefghijklmnop", "ponmlkjihgfedcba"}},
		},
		{
			name:  "Numbers as strings",
			input: []string{"123", "321", "213", "456"},
			want:  [][]string{{"123", "213", "321"}, {"456"}},
		},
		{
			name:  "Special characters",
			input: []string{"a!b", "b!a", "a?b"},
			want:  [][]string{{"a!b", "b!a"}, {"a?b"}},
		},
		{
			name:  "Empty array",
			input: []string{},
			want:  [][]string{},
		},
	}

	for name, solution := range solutions {
		t.Run(name, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					got := solution(tc.input)
					normalizedGot := normalizeGroups(got)
					normalizedWant := normalizeGroups(tc.want)
					if !reflect.DeepEqual(normalizedGot, normalizedWant) {
						t.Errorf("GroupAnagrams(%v) = %v; want %v", tc.input, got, tc.want)
					}
				})
			}
		})
	}
}
