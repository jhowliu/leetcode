package SquareSum

import "math"

func judgeSquareSum(c int) bool {
	hashTable := map[int]int{}

	upper := int(math.Sqrt(float64(c)))

	for i := 0; i <= upper; i++ {
		square := i * i
		if _, ok := hashTable[square]; ok {
			return true
		}

		if c-square == square {
			return true
		}

		hashTable[c-square] = 1
	}

	return false
}
