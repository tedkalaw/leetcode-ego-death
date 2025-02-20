package solutions

// You are given an integer n representing the number of steps to reach the top of a staircase.
// You can climb with either 1 or 2 steps at a time.

// Return the number of distinct ways to climb to the top of the staircase.
func ClimbingStairs(n int) int {
	if n <= 2 {
		return n
	}

	// dp := []int{}
	// The second arg of make for slices is length
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
