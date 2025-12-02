package main

import "fmt"

func main() {
	fmt.Println("Permutations", permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	res := [][]int{}
	helper(0, nums, &res)
	return res
}

func helper(ind int, nums []int, res *[][]int) {
	fmt.Printf("helper called with ind:%d nums:%v\n", ind, nums)
	if ind == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		fmt.Printf("res updated:%v\n", *res)
		return
	}

	for i := ind; i < len(nums); i++ {
		fmt.Printf("i: %d, swapping %d and %d\n", i, nums[ind], nums[i])
		nums[ind], nums[i] = nums[i], nums[ind]
		helper(ind+1, nums, res)
		fmt.Printf("i: %d, backtracking by swapping %d and %d\n", i, nums[ind], nums[i])
		nums[ind], nums[i] = nums[i], nums[ind]
	}
}

// n to 0 code

func permuteNto0(nums []int) [][]int {
	res := [][]int{}
	helperNto0(len(nums)-1, nums, &res)
	return res
}

func helperNto0(ind int, nums []int, res *[][]int) {
	if ind < 0 {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}

	for i := ind; i >= 0; i-- {
		nums[ind], nums[i] = nums[i], nums[ind]
		helperNto0(ind-1, nums, res)
		nums[ind], nums[i] = nums[i], nums[ind]
	}

}
