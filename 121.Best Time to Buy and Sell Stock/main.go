/*
ay you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.

Example 1:

Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
             Not 7-1 = 6, as selling price needs to be larger than buying price.
Example 2:

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
*/
package main

import (
	"log"
	"math"
)

// 動態規劃問題 Maximum Subarray problem
// A: Kadane's Alogorithm: O(n)內
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	maxProfit := 0
	curProfit := 0

	for i := 1; i != len(prices); i++ {
		// curProfit+diff (tricky) 加總後結果跟0比較,因為目標是求集合內最大值,故小於0不考慮(小於0代表數字i更小, 更小表示可能有更大利益從i開始)
		// 若加總結果較大則表示不需要重新計算,若0較大則代表加總需要重新由0開始計算(代表前一個集合無法得到最大值)
		diff := prices[i] - prices[i-1]
		curProfit = int(math.Max(0.0, float64(curProfit+diff)))
		// 記錄著一路上碰到的最大加總(各種可能的數字集合結果)
		maxProfit = int(math.Max(float64(maxProfit), float64(curProfit)))
	}

	return maxProfit
}

func main() {
	testCases := []struct {
		description string
		input       []int
		expect      int
	}{
		{
			description: "Test Case 1",
			input:       []int{7, 1, 5, 3, 6, 4},
			expect:      5,
		},
		{
			description: "Test Case 2",
			input:       []int{7, 6, 4, 3, 1},
			expect:      0,
		},
	}

	for _, t := range testCases {
		result := maxProfit(t.input)
		if result != t.expect {
			log.Fatalf(
				"%s: expect[%v] != actual[%v]",
				t.description, t.expect, result,
			)
		}
	}
}
