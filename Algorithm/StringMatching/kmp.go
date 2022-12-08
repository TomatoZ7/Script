package StringMatching

func getNexts(s string, lenS int) []int {
	next := make([]int, lenS)
	next[0] = -1

	k := -1
	for i := 1; i < lenS; i++ {
		for k != -1 && s[k+1] != s[k] {
			k = next[k]
		}
		if s[k+1] == s[i] {
			k++
		}
		next[i] = k
	}

	return next
}

func KMP(ss, s string) int {
	lenSS, lenS := len(ss), len(s)
	next := getNexts(s, lenS)

	j := 0
	for i := 0; i < lenSS; i++ {
		for j > 0 && ss[i] != s[j] {
			j = next[j-1]+1
		}
		if ss[i] == s[j] {
			j++
		}
		// 模式串匹配成功
		if j == lenS {
			return i - lenS + 1
		}
	}

	return -1
}