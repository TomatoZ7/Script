package NQueens

import "math"

func ViolentN(row int) {
	// 全部 Queen 摆放完成，判断是否符合逻辑
	if n == row {
		isLegal := true

		for i := 0; i < n; i++ {
			for j := i+1; j < n; j++ {
				if math.Abs(float64(i - j)) == math.Abs(float64(board[i] - board[j])) {
					isLegal = false
					return
				}
			}
		}

		if isLegal {
			res++
		}

		return
	}

	// 循环每列，放置 Queen
	for col := 0; col < n; col++ {
		if !queenPutMap[col] {		// 值为 false 说明可以放入
			board[row] = col		// 将 Queen 放入第 row 行第 col 列
			queenPutMap[col] = true
			ViolentN(row + 1)
			queenPutMap[col] = false
		}
	}
}