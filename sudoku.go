package main

import (
	"fmt"
)

func printBoard(board [9][9]int) {
	for _, row := range board {
		for _, num := range row {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}

func isValid(board [9][9]int, row, col, num int) bool {
	// Satırı kontrol et
	for x := 0; x < 9; x++ {
		if board[row][x] == num {
			return false
		}
	}

	// Sütunu kontrol et
	for x := 0; x < 9; x++ {
		if board[x][col] == num {
			return false
		}
	}

	// 3x3 kutuyu kontrol et
	startRow := row - row%3
	startCol := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i+startRow][j+startCol] == num {
				return false
			}
		}
	}
	return true
}

func solveSudoku(board *[9][9]int) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == 0 { // Boş hücre bul
				for num := 1; num <= 9; num++ { // 1-9 arası sayıları dene
					if isValid(*board, row, col, num) {
						board[row][col] = num
						if solveSudoku(board) {
							return true
						}
						board[row][col] = 0 // Geri dön
					}
				}
				return false // Geri dönmeyi tetikle
			}
		}
	}
	return true // Bulmaca çözüldü
}

func main() {
	// Örnek Sudoku tahtası (0 boş hücreleri temsil eder)
	sudokuBoard := [9][9]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	fmt.Println("Sudoku Tahtası:")
	printBoard(sudokuBoard)

	if solveSudoku(&sudokuBoard) {
		fmt.Println("\nÇözülmüş Sudoku Tahtası:")
		printBoard(sudokuBoard)
	} else {
		fmt.Println("Çözüm yok.")
	}
}
