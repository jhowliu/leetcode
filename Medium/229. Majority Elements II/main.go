package solution

func findMajorityElements(nums []int) []int {
	n := len(nums)
	result := []int{}

	if n == 0 {
		return result
	}

	cntA, cntB, candidateA, candidateB := 0, 0, -1, -2

	for _, num := range nums {
		if candidateA == num {
			cntA++
		} else if candidateB == num {
			cntB++
		} else if cntA == 0 {
			cntA, candidateA = 1, num
		} else if cntB == 0 {
			cntB, candidateB = 1, num
		} else {
			cntA, cntB = cntA-1, cntB-1
		}
	}

	cntA, cntB = 0, 0

	// re-calculate count
	for _, num := range nums {
		if num == candidateA {
			cntA++
		}

		if num == candidateB {
			cntB++
		}
	}

	if cntA > n/3 {
		result = append(result, candidateA)
	}
	if cntB > n/3 {
		result = append(result, candidateB)
	}

	return result
}
