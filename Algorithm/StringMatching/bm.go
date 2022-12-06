package StringMatching

// 模式串数组的大小
const modelArraySize = 256
// 模式串数组，用于记录每个字符在模式串中出现的最后位置
var modelArr = [modelArraySize]int{}

// 记录模式串中每个字符最后出现的位置
func generateModelArr(s string) {
	// 初始化数组
	for i := 0; i < len(modelArr); i++ {
		modelArr[i] = -1
	}
	// 赋值
	for i, ch := range s {
		modelArr[ch] = i
	}
}

// 生成好后缀 suffix 和 prefix
func generateGS(s string, suffix []int, prefix []bool) {
	lenS := len(s)
	// 初始化
	for i := 0; i < lenS; i++ {
		suffix[i] = -1
		prefix[i] = false
	}

	for i := 0; i < lenS - 1; i++ {
		j := i
		// 公共后缀子串长度
		k := 0

		for j >= 0 && s[j] == s[lenS-1-k] {
			j--
			k++
			// j+1 表示公共后缀子串在 s[0, i] 中的起始下标
			suffix[k] = j + 1
		}
		// 如果公共后缀子串也是模式串的前缀子串
		if j == -1 {
			prefix[k] = true
		}
	}
}

// 计算 好后缀 情况下的偏移量
// j 是指坏字符的字符下标
func calOffsetByGS(j, lenS int, suffix []int, prefix []bool) int {
	// 好后缀长度
	k := lenS - 1 - j
	if suffix[k] != -1 {
		return j - suffix[k] + 1
	}

	for r := j + 2; r <= lenS - 1; r++ {
		if prefix[lenS-r] {
			return r
		}
	}

	return lenS
}

// BM 算法
func BM(ss, s string) int {
	lenSS, lenS := len(ss), len(s)
	// init: 坏字符
	generateModelArr(s)
	// init: 好后缀s
	suffix := make([]int, lenS)
	prefix := make([]bool, lenS)
	generateGS(s, suffix, prefix)

	i := 0
	for i <= lenSS - lenS {
		var j int

		// 模式串从后往前匹配
		for j = lenS - 1; j >= 0; j-- {
			if ss[i+j] != s[j] {	// 此时坏字符对应的下标为 j
				break
			}
		}

		// 匹配成功，返回主串下标 i
		if j < 0 {
			return i
		}

		// 坏字符条件下的偏移量
		offsetBC := j - modelArr[ss[i+j]]

		// 好后缀条件下的偏移量
		offsetGS := 0
		if j < lenS - 1 {
			offsetGS = calOffsetByGS(j, lenS, suffix, prefix)
		}

		// 这里等同于将模式串往后滑动 j - modelArr[ss[i+j]] 位
		if offsetBC > offsetGS {
			i += offsetBC
		} else {
			i += offsetGS
		}

	}

	return -1
}
