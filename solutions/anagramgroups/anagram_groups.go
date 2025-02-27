package anagramgroups

import (
	"slices"
)

func SortingSolution(strs []string) [][]string {
	sortedStringMap := make(map[string][]string)

	for _, s := range strs {

		runes := []rune(s)
		slices.Sort(runes)
		// string can convert a slice of runes back into a string; below
		// I tried to use strings.Join, but that doesn't work because
		// strings.Join takes a slice of strings, not runes
		sortedStr := string(runes)

		// sortedStringMap[strings.Join(runes, "")] = s
		// We need to use the append function here
		sortedStringMap[sortedStr] = append(sortedStringMap[sortedStr], s)
	}

	result := [][]string{}
	for _, anagramGroup := range sortedStringMap {
		result = append(result, anagramGroup)
	}

	return result
}
