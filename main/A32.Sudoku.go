package main

import (
	"fmt"
)

func main() {
	board := [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}

	valid := validateSudokuBoard(board)

	if valid {
		fmt.Println("O tabuleiro de Sudoku é válido.")
	} else {
		fmt.Println("O tabuleiro de Sudoku contém células com valores incorretos.")
	}
}

func validateSudokuBoard(board [][]int) bool {
	for _, row := range board {
		if !isValidSet(row) {
			return false
		}
	}

	for col := 0; col < 9; col++ {
		column := []int{}
		for row := 0; row < 9; row++ {
			column = append(column, board[row][col])
		}
		if !isValidSet(column) {
			return false
		}
	}

	for startRow := 0; startRow < 9; startRow += 3 {
		for startCol := 0; startCol < 9; startCol += 3 {
			region := []int{}
			for row := startRow; row < startRow+3; row++ {
				for col := startCol; col < startCol+3; col++ {
					region = append(region, board[row][col])
				}
			}
			if !isValidSet(region) {
				return false
			}
		}
	}

	return true
}

func isValidSet(nums []int) bool {
	visited := make(map[int]bool)
	for _, num := range nums {
		if num != 0 {
			if visited[num] {
				return false
			}
			visited[num] = true
		}
	}
	return true
}