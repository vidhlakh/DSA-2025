package main

import "fmt"

func main() {
	fmt.Println("Num Decodings", numDecodings("226"))
}

var (
	alpMap = map[string]string{
		"1":  "A",
		"2":  "B",
		"3":  "C",
		"4":  "D",
		"5":  "E",
		"6":  "F",
		"7":  "G",
		"8":  "H",
		"9":  "I",
		"10": "J",
		"11": "K",
		"12": "L",
		"13": "M",
		"14": "N",
		"15": "O",
		"16": "P",
		"17": "Q",
		"18": "R",
		"19": "S",
		"20": "T",
		"21": "U",
		"22": "V",
		"23": "W",
		"24": "X",
		"25": "Y",
		"26": "Z",
	}
)

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	if s[0] == '0' {
		return 0
	}
	var res []string
	cc := helper(0, s, len(s), "", &res)
	fmt.Printf("result:%v", res)
	return cc
}

func helper(ind int, s string, n int, curr string, res *[]string) int {

	// condition
	if ind == n {

		*res = append(*res, curr)
		return 1
	}
	// having 0 at start is invalid eg. 0 for single digit 06 for two digit
	if s[ind] == '0' {
		return 0
	}

	//single digit

	count := helper(ind+1, s, n, curr+alpMap[string(s[ind])], res)
	//two digit
	if ind+1 < n && (s[ind] == '1' || (s[ind] == '2' && s[ind+1] <= '6')) {
		count += helper(ind+2, s, n, curr+alpMap[fmt.Sprintf("%c%c", s[ind], s[ind+1])], res)
	}
	return count
}
