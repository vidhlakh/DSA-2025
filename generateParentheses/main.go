package main

import "fmt"

// https://leetcode.com/problems/generate-parentheses/

func main() {
	fmt.Println("Generate paranthesesis", generateParenthesis(3))
}
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	helper(0, 0, n, "", &res)
	return res
}

func helper(numOfOpen, numOfClose, n int, curr string, res *[]string) {

	// condn
	if numOfClose == n {
		*res = append(*res, curr)

	}
	// open
	if numOfOpen < n {
		helper(numOfOpen+1, numOfClose, n, curr+"(", res)

	}

	// close
	if numOfOpen > numOfClose {
		helper(numOfOpen, numOfClose+1, n, curr+")", res)
	}

}
