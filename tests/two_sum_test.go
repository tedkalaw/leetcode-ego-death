package tests

import (
	"reflect"
	"testing"

	"github.com/tedkalaw/leetcode-ego-death/solutions/two_sum"
)

func TestTwoSum(t *testing.T) {
	solutions := map[string]func([]int, int) []int{
		"hashmap": two_sum.HashMapSolution,
		"brute":   two_sum.BruteForceSolution,
	}

	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"Example1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"Example2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"NoSolution", []int{5, 5, 5}, 10, []int{0, 1}}, // Adjust as needed
	}

	for name, solution := range solutions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := solution(tt.nums, tt.target)
					if !reflect.DeepEqual(got, tt.want) {
						t.Errorf("TwoSum() = %v, want %v", got, tt.want)
					}
				})
			}
		})
	}
}
