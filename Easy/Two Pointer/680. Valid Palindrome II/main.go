package Palindrome

func validPalindrome(s string) bool {
	i, j := 0, len(s)-1

	for i < j {
		if s[i] != s[j] {
			// 當s[i+1]==s[j] && s[i]==s[j-1]，刪除會影響後面結果，所以兩個都需要嘗試。
			return isPalindrome(i+1, j, s) || isPalindrome(i, j-1, s)
		}
		i++
		j--
	}

	return true
}

func isPalindrome(i int, j int, s string) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
