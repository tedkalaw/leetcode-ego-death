package solutions

// ValidAnagram determines if two strings are anagrams of each other.
// An anagram is a string that contains the exact same characters as another string,
// but the order of the characters can be different.
// The strings consist of lowercase English letters only.
func ValidAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	freqMap := make(map[rune]int)
	for _, r := range s {
		freqMap[r]++
	}

	for _, r := range t {
		freqMap[r]--
		if freqMap[r] < 0 {
			return false
		}
	}

	return true
}
