package NQueens

// 棋盘大小
var n = 4
// 棋盘，下标表示行，值表示该行的第几列放了 Queen
var board []int
// 最终结果
var res int

// 在暴力破解中用于记录哪一列已经摆放了 Queen
var hashTable []bool
var queenPutMap []bool