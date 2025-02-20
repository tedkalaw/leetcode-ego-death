package solutions

// import "fmt"

// Given an integer array nums, return true if any value appears at least
// twice in the array, and return false if every element is distinct.
func ContainsDuplicate(nums []int) bool {
	// seenSet := make(map[int]bool)
	// Apparently for a set like this, using an empty struct as the value type
	// uses 0 bytes
	// Seems rather sweaty
	seenSet := make(map[int]struct{})

	for _, n := range nums {
		// The short statement before the condition in Go is kinda cool
		// Curious to see how common that is
		if _, exists := seenSet[n]; exists {
			return true
		}
		seenSet[n] = struct{}{}
	}

	return false
}
