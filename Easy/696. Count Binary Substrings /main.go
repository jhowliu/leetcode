package solution

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func countBinarySubstrings(s string) int {
	if len(s) == 0 {
		return 1
	}

	pre, cur, count := 0, 1, 0

	for i := 1; i < len(s); i++ {
		if int(s[i]) == int(s[i-1]) {
			cur++
		} else {
			// 到這裡代表遇到不同的數字,然後跟上一個數字的長度比較，最低的代表substring數量
			// Ex: 0001111 => pre, cur => 3, 4 => min(3, 4) => 3 (最多3個substring)
			count += min(pre, cur)
			pre = cur
			cur = 1
		}
	}

	// 最後一次要記得比較
	return count + min(pre, cur)
}
