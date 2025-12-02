package main

import (
	"fmt"
	"strings"
)

func main() {
	solutions := solveNQueens(4)

	// rowwise approach
	for si, sol := range solutions {
		fmt.Printf("Solution %d:\n", si+1)
		for _, row := range sol {
			fmt.Println(row)
		}
		fmt.Println()
	}

	// colwise approach
	solutionsColwise := solveNQueensColwise(4)
	for si, sol := range solutionsColwise {
		fmt.Printf("Colwise Solution %d:\n", si+1)
		for _, row := range sol {
			fmt.Println(row)
		}
		fmt.Println()
	}
}

func solveNQueens(n int) [][]string {
	if n <= 0 {
		return nil
	}

	// board as mutable bytes per row for easy updates
	board := make([][]byte, n)
	rowTemplate := strings.Repeat(".", n)
	for i := 0; i < n; i++ {
		board[i] = []byte(rowTemplate)
	}

	res := [][]string{}
	helperRowwise(0, board, &res, n)
	return res
}

func solveNQueensColwise(n int) [][]string {
	if n <= 0 {
		return nil
	}

	// board as mutable bytes per row for easy updates
	board := make([][]byte, n)
	rowTemplate := strings.Repeat(".", n)
	for i := 0; i < n; i++ {
		board[i] = []byte(rowTemplate)
	}

	res := [][]string{}
	helperColwise(0, board, &res, n)
	return res
}

func helperRowwise(row int, board [][]byte, res *[][]string, n int) {
	if row == n {
		// convert board to []string and append
		sol := make([]string, n)
		for i := 0; i < n; i++ {
			sol[i] = string(board[i])
		}
		*res = append(*res, sol)
		return
	}

	for col := 0; col < n; col++ {
		if isValidForRowwiseInsertion(board, row, col, n) {
			board[row][col] = 'Q'
			helperRowwise(row+1, board, res, n)
			board[row][col] = '.'
		}
	}
}

func isValidForRowwiseInsertion(board [][]byte, row, col, n int) bool {
	// check vertical up
	for r := 0; r < row; r++ {
		if board[r][col] == 'Q' {
			return false
		}
	}
	// check left-up diagonal
	for r, c := row-1, col-1; r >= 0 && c >= 0; r, c = r-1, c-1 {
		if board[r][c] == 'Q' {
			return false
		}
	}
	// check right-up diagonal
	for r, c := row-1, col+1; r >= 0 && c < n; r, c = r-1, c+1 {
		if board[r][c] == 'Q' {
			return false
		}
	}
	return true
}

// helper colwise approach
func helperColwise(col int, board [][]byte, res *[][]string, n int) {
	if col == n {
		// convert board to []string and append
		sol := make([]string, n)
		for i := 0; i < n; i++ {
			sol[i] = string(board[i])
		}
		*res = append(*res, sol)
		return
	}

	for row := 0; row < n; row++ {
		if isValidForColwiseInsertion(board, row, col, n) {
			board[row][col] = 'Q'
			helperColwise(col+1, board, res, n)
			board[row][col] = '.'
		}
	}
}

func isValidForColwiseInsertion(board [][]byte, row, col, n int) bool {
	// check horizontal left
	for c := 0; c < col; c++ {
		if board[row][c] == 'Q' {
			return false
		}
	}
	// check left-up diagonal
	for r, c := row-1, col-1; r >= 0 && c >= 0; r, c = r-1, c-1 {
		if board[r][c] == 'Q' {
			return false
		}
	}
	// check left-down diagonal
	for r, c := row+1, col-1; r < n && c >= 0; r, c = r+1, c-1 {
		if board[r][c] == 'Q' {
			return false
		}
	}
	return true
}
