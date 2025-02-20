package tests

import (
	"testing"

	"github.com/tedkalaw/leetcode-ego-death/solutions"
)

func TestClimbingStairs(t *testing.T) {
	testCases := []struct {
		name  string
		input int
		want  int
	}{
		{
			name:  "n=1",
			input: 1,
			want:  1,
		},
		{
			name:  "n=2",
			input: 2,
			want:  2,
		},
		{
			name:  "n=3",
			input: 3,
			want:  3,
		},
		{
			name:  "n=4",
			input: 4,
			want:  5,
		},
		{
			name:  "n=5",
			input: 5,
			want:  8,
		},
		{
			name:  "n=6",
			input: 6,
			want:  13,
		},
		{
			name:  "n=7",
			input: 7,
			want:  21,
		},
		{
			name:  "n=8",
			input: 8,
			want:  34,
		},
		{
			name:  "n=9",
			input: 9,
			want:  55,
		},
		{
			name:  "n=10",
			input: 10,
			want:  89,
		},
		{
			name:  "n=11",
			input: 11,
			want:  144,
		},
		{
			name:  "n=12",
			input: 12,
			want:  233,
		},
		{
			name:  "n=13",
			input: 13,
			want:  377,
		},
		{
			name:  "n=14",
			input: 14,
			want:  610,
		},
		{
			name:  "n=15",
			input: 15,
			want:  987,
		},
		{
			name:  "n=16",
			input: 16,
			want:  1597,
		},
		{
			name:  "n=20",
			input: 20,
			want:  10946,
		},
		{
			name:  "n=25",
			input: 25,
			want:  121393,
		},
		{
			name:  "n=28",
			input: 28,
			want:  514229,
		},
		{
			name:  "n=29",
			input: 29,
			want:  832040,
		},
		{
			name:  "n=30",
			input: 30,
			want:  1346269,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := solutions.ClimbingStairs(tc.input)
			if got != tc.want {
				t.Errorf("ClimbingStairs(%v) = %v; want %v", tc.input, got, tc.want)
			}
		})
	}
}
