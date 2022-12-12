package NQueens

// 棋盘大小
var n = 8
// 棋盘，下标表示行，值表示该行的第几列放了 Queen
var board []int
// 最终结果
var res int

// 在暴力破解中用于记录哪一列已经摆放了 Queen
var queenPutMap []bool

func init() {
	board = make([]int, n+2)
	queenPutMap = make([]bool, n+2)	// 下标代表列，值代表该列是否存在 Queen
}