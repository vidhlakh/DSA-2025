package main

import "fmt"

func main() {
	fmt.Println("Min cost of climbing stairs", minCostClimbingStairs([]int{10, 15, 20}))
}
func minCostClimbingStairs(cost []int) int {
	return min(helper(0, cost), helper(1, cost))
}

func helper(ind int, cost []int) int {
	if ind >= len(cost) {
		return 0
	}
	//fmt.Println("ind",ind)
	return cost[ind] + min(helper(ind+2, cost), helper(ind+1, cost))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
