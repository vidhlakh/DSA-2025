package main

import "fmt"

func main() {
	fibTopDown()
	fibBottomUp()
	fibBottomUpSpaceOptimized()
}

// top down approach
func fibTopDown() {
	n := 10                // Example input
	dp := make([]int, n+1) // its n+1 because we want to store from 0 to n
	for i := range len(dp) {
		dp[i] = -1
	}
	result := helperTopDown(n, dp)
	fmt.Println("Fibonacci of", n, "is", result)
}

// Top Down approach Bigger problem broken into smaller subproblems and store the results of subproblems to avoid recomputation
// Recursion is always used in Top Down approach
func helperTopDown(n int, dp []int) int {

	if dp[n] != -1 {
		fmt.Println("Using memoized value for n =", n)
		return dp[n]
	}
	fmt.Println("Calculating fib(", n, ")")
	if n <= 1 {
		fmt.Println("Base case reached for n =", n)
		return n
	}

	dp[n] = helperTopDown(n-1, dp) + helperTopDown(n-2, dp)
	return dp[n]
}

func fibBottomUp() {
	n := 10                // Example input
	dp := make([]int, n+1) // its n+1 because we want to store from 0 to n

	result := helperBottomUp(n, dp)
	fmt.Println("Fibonacci of", n, "is", result)
}

// Bottom Up approach means Smaller subproblems are solved first and their results are used to build up solutions to bigger problems
// Iteration is always used in Bottom Up approach
func helperBottomUp(n int, dp []int) int {
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func fibBottomUpSpaceOptimized() {
	n := 10 // Example input

	result := helperBottomUpSpaceOptimized(n)
	fmt.Println("Fibonacci of", n, "is", result)
}

func helperBottomUpSpaceOptimized(n int) int {
	if n <= 1 {
		return n
	}
	prev2 := 0
	prev1 := 1
	var curr int
	for i := 2; i <= n; i++ {
		curr = prev1 + prev2
		prev2 = prev1
		prev1 = curr
	}
	return curr

}
