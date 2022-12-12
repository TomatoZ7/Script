package main

import (
	"fmt"
)

var n = 4
var board = []int{0, 3, -1, -1}

func main() {
	//NQueens.CalQueens()
	fmt.Println(putQueen(2, 1))
}

// 放置 Queen
func putQueen(row, column int) bool {

	for ; column < n; column++ {
		if canPutQueen(row, column) {
			// 可以摆放
			board[row] = column
			return true
		}
	}

	return false
}

// 判断第 row 行，第 column 列能否放置
func canPutQueen(row, column int) bool {
	// 取前后列的值，方便计算「左上角」和「右上角」
	leftUp, rightUp := column - 1, column + 1

	// 逐行往上判断每一行
	for i := row - 1; i >= 0; i-- {
		// 如果上一行就摆放了 Queen，则为 False
		if board[i] == column {
			return false
		}

		// 如果左上对角线摆放了 Queen，则为 False
		if leftUp >= 0 && board[i] == leftUp {
			return false
		}

		// 如果右上对角线摆放了 Queen，则为 False
		if rightUp < n && board[i] == rightUp {
			return false
		}
	}

	return true
}

func init() {
	// 初始化 board
	for i := 0; i < len(board); i++ {
		board[i] = -1
	}
}