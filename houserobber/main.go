package main

import (
	"fmt"
)

// house robber
// https://leetcode.com/problems/house-robber/description/
// house robber II - circular houses
// https://leetcode.com/problems/house-robber-ii/

func main() {

	res1 := houserobber1([]int{2, 7, 9, 3, 1})
	fmt.Println("House robber I:", res1)
	res2 := houserobber2([]int{2, 3, 2})
	fmt.Println("House robber II:", res2)
}
func houserobber2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		max(nums[0], nums[1])
	}
	return max(helper(0, nums, len(nums)-2), helper(1, nums, len(nums)-1))
}

func helper(ind int, nums []int, length int) int {

	if ind > length {
		return 0
	}

	rob := nums[ind] + helper(ind+2, nums, length)
	dontrob := helper(ind+1, nums, length)
	return max(rob, dontrob)

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func houserobber1(nums []int) int {
	return helper1(0, nums)
}

func helper1(ind int, nums []int) int {
	if ind >= len(nums) {
		return 0
	}
	rob := nums[ind] + helper1(ind+2, nums)
	dontrob := helper1(ind+1, nums)
	return max(rob, dontrob)

}
