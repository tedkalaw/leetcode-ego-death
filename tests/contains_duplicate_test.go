package tests

import (
	"testing"

	"github.com/tedkalaw/leetcode-ego-death/solutions"
)

func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		want  bool
	}{
		{
			name:  "No duplicates",
			input: []int{1, 2, 3, 4},
			want:  false,
		},
		{
			name:  "Has duplicates",
			input: []int{1, 2, 3, 1},
			want:  true,
		},
		{
			name:  "Single element",
			input: []int{10},
			want:  false,
		},
		{
			name:  "Multiple duplicates",
			input: []int{1, 1, 2, 3, 2},
			want:  true,
		},
		{
			name:  "All same",
			input: []int{5, 5, 5, 5},
			want:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := solutions.ContainsDuplicate(tc.input)
			if got != tc.want {
				t.Errorf("ContainsDuplicate(%v) = %v; want %v", tc.input, got, tc.want)
			}
		})
	}
}
