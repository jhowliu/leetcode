package solution

func longestPalindrome(s string) int {
	counter := map[int]int{}

	for _, c := range s {
		counter[int(c)]++
	}

	result := 0
	hadOdd := false

	for _, v := range counter {
		// 回文其實就是字母出現的兩倍
		result += (v / 2) * 2

		// 若找到單數的代表可以再多一長度
		if v%2 == 1 {
			hadOdd = true
		}
	}

	if hadOdd {
		return result + 1
	}

	return result
}
