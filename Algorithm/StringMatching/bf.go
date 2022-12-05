package StringMatching

func BF(str1, str2 string) int {
	len1, len2 := len(str1), len(str2)
	if len2 > len1 {
		return -1
	}

	for i1 := 0; i1 < len1; i1++ {
		for i2 := 0; i2 < len2 && i1 + i2 < len1; i2++ {
			if str2[i2] != str1[i1+i2] {
				break
			}

			if i2 == len2 - 1 {
				return i1
			}
		}
	}

	return -1
}