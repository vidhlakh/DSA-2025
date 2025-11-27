package main

import "fmt"

func main() {
	// Example usage
	fmt.Println("combination possible:", combine(4, 2))
	fmt.Println("combination possible:", combine2(4, 2))
}

// 1 to n
func combine(n int, k int) [][]int {
	res := [][]int{}
	helper(1, n, k, []int{}, &res)
	return res
}

func helper(ind, n, k int, curr []int, res *[][]int) {
	// condition
	if k == 0 {
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		*res = append(*res, tmp)
		return
	}
	if ind > n {
		return
	}
	//include
	curr = append(curr, ind)
	helper(ind+1, n, k-1, curr, res)
	curr = curr[:len(curr)-1]
	//exclude
	helper(ind+1, n, k, curr, res)

}

// n to 1
func combine2(n int, k int) [][]int {
	res := [][]int{}
	helper2(n, n, k, []int{}, &res)
	return res
}

func helper2(ind, n, k int, curr []int, res *[][]int) {
	// condition
	if k == 0 {
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		*res = append(*res, tmp)
		return
	}
	if ind < 1 {
		return
	}
	//include
	curr = append(curr, ind)
	helper(ind-1, n, k-1, curr, res)
	curr = curr[:len(curr)-1]
	//exclude
	helper(ind-1, n, k, curr, res)

}
