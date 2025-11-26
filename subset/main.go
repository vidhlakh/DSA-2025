package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{3, 2, 4, 1}))
}
func subsets(nums []int) [][]int {
	res := [][]int{}
	helper(0, []int{}, nums, &res)
	return res
}

func helper(ind int, curr []int, nums []int, res *[][]int) {
	if ind >= len(nums) {
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		*res = append(*res, tmp)
		return
	}
	// include
	curr = append(curr, nums[ind])
	helper(ind+1, curr, nums, res)
	curr = curr[:len(curr)-1]

	//exclude
	helper(ind+1, curr, nums, res)

}
