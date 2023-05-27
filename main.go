package main

import "math"

func main() {
	// put your sudoku board in here
	board := [][]int8{
		{0, 9, 7, 0, 8, 0, 0, 0, 4},
		{0, 0, 0, 0, 0, 7, 1, 0, 0},
		{3, 0, 2, 0, 0, 0, 0, 6, 0},
		{0, 0, 9, 0, 0, 0, 0, 0, 0},
		{6, 0, 0, 1, 0, 2, 0, 0, 0},
		{0, 3, 0, 5, 9, 0, 0, 0, 2},
		{0, 0, 0, 8, 7, 0, 0, 3, 5},
		{0, 0, 3, 2, 0, 6, 9, 0, 0},
		{8, 5, 0, 3, 0, 0, 0, 0, 0},
	}

	game := createBoardGame(int(math.Sqrt(float64(len(board)))))

	for i := range board {
		for j, val := range board[i] {
			if val > 0 {
				game.addFixedNumber(val, i, j)
			}
		}
	}

	proccess(&game)
	// see the output text for the result
}
