package solution

func isAnagram(s, t string) bool {
	counter := make([]int, 26)

	for _, c := range s {
		counter[c-'a']++
	}

	for _, c := range t {
		counter[c-'a']--
	}

	for _, v := range counter {
		if v != 0 {
			return false
		}
	}

	return true
}
