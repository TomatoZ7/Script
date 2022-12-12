package NQueens

import (
	"fmt"
	"math"
)

func BackTracking(row int) {
	if row == n {
		res++
		return
	}

	for col := 0; col < n; col++ {
		// !queenPutMap[col] 用来控制列，为 True 则说明该列已经存在 Queen
		if !queenPutMap[col] {
			isLegal := true
			for preRow := 0; preRow < row; preRow++ {
				if math.Abs(float64(row - preRow)) == math.Abs(float64(col - board[preRow])) {
					isLegal = false
					break
				}
			}

			if isLegal {
				board[row] = col
				queenPutMap[col] = true
				BackTracking(row + 1)
				queenPutMap[col] = false
			}
		}
	}
}

func PrintRes() {
	fmt.Println(res)
}