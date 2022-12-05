package StringMatching

import (
	"crypto/md5"
	"encoding/hex"
)

func RK(str1, str2 string) int {
	len1, len2 := len(str1), len(str2)
	if len2 > len1 {
		return -1
	}

	// 用于存放子串 MD5
	var s1 []string
	// 计算模式串的 MD5
	m := md5.New()
	m.Write([]byte(str2))
	str2Md5 := hex.EncodeToString(m.Sum(nil))

	// 生成子串 MD5
	for i := 0; i <= len1 - len2; i++ {
		m = md5.New()
		m.Write([]byte(str1[i:i+len2]))
		s1 = append(s1, hex.EncodeToString(m.Sum(nil)))
	}

	for i, strMd5 := range s1 {
		if strMd5 == str2Md5 {
			return i
		}
	}

	return -1
}
