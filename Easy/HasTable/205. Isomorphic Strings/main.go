package solution

func isIsomorphic(s string, t string) bool {
	preIndexOfS := map[int]int{}
	preIndexOfT := map[int]int{}

	for i := range s {
		if preIndexOfS[int(s[i])] != preIndexOfT[int(t[i])] {
			return false
		}

		// + 1是為了避免沒出現過跟出現在index=0的位置混淆 ex: aa, ab
		preIndexOfS[int(s[i])] = i + 1
		preIndexOfT[int(t[i])] = i + 1
	}

	return true
}
