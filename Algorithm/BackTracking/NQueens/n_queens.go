package NQueens

import "fmt"

/*
 * N 皇后算法：在指定的 N 乘 N 棋盘内，要求摆放 N 个皇后，使得每个皇后所在的列、行、斜线都只能存在一个皇后，求一共有多少种摆法。
 */

// 棋盘大小
var n = 4
// 棋盘，下标表示行，值表示该行的第几列放了 Queen
var board []int

func init() {
	// 初始化 board
	board = make([]int, n)
	for i := 0; i < len(board); i++ {
		board[i] = -1
	}
}

// 放置 Queen
func putQueen(row, column int) bool {

	for ; column < n; column++ {
		if row == 2 {
			fmt.Println("", canPutQueen(row, column), row, column)
		}
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

func CalQueens() {
	// 从第 1 行开始摆放
	r, c := 0, 0
	for r < n {
		if r == 2 {
			fmt.Println("第2行：", r, c, board)
		}
		res := putQueen(r, c)

		// 所有行成功摆放
		//if res && r == n - 1 {
		//	fmt.Println(res)
		//}

		if r == 2 {
			fmt.Println("结果：", res, board)
			fmt.Println("------")
		}

		// 摆放失败，则回溯
		if !res {
			c = board[r-1] + 1
			board[r-1] = -1
			r -= 1
		} else {
			c = 0
			r++
		}
	}

	fmt.Println(board)
}
