package two_sum

func HashMapSolution(nums []int, target int) []int {
	indices := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, found := indices[complement]; found {
			return []int{j, i}
		}
		indices[num] = i
	}
	return []int{}
}

func BruteForceSolution(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
