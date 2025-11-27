package main

import "fmt"

func main() {
	fmt.Println("Num Decodings", numDecodings("27306"))
}
func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	if s[0] == '0' {
		return 0
	}
	return helper(0, s, len(s))
}

func helper(ind int, s string, n int) int {

	// condition
	if ind == n {
		return 1
	}
	// having 0 at start is invalid eg. 0 for single digit 06 for two digit
	if s[ind] == '0' {
		return 0
	}

	//single digit
	count := helper(ind+1, s, n)
	//two digit
	if ind+1 < n && (s[ind] == '1' || (s[ind] == '2' && s[ind+1] <= '6')) {
		count += helper(ind+2, s, n)
	}
	return count
}
