package solutions

func TwoSum(nums []int, target int) []int {
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
