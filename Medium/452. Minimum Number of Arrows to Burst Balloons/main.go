/*

There are a number of spherical balloons spread in two-dimensional space. For each balloon, provided input is the start and end coordinates of the horizontal diameter. Since it's horizontal, y-coordinates don't matter and hence the x-coordinates of start and end of the diameter suffice. Start is always smaller than end. There will be at most 104 balloons.

An arrow can be shot up exactly vertically from different points along the x-axis. A balloon with xstart and xend bursts by an arrow shot at x if xstart ≤ x ≤ xend. There is no limit to the number of arrows that can be shot. An arrow once shot keeps travelling up infinitely. The problem is to find the minimum number of arrows that must be shot to burst all balloons.


Example:

	Input:
		[[10,16], [2,8], [1,6], [7,12]]

	Output:
		2

	Explanation:
		One way is to shoot one arrow for example at x = 6 (bursting the balloons [2,8] and [1,6]) and another arrow at x = 11 (bursting the other two balloons).

問題： 找出花費最少的飛鏢數量，射破所有氣球。(重疊的區間代表只需要一支飛鏢，所以有多少個非重疊區域就是飛鏢數量)
*/
package main

import (
	"sort"
)

type Points [][]int

func (slice Points) Len() int {
	return len(slice)
}

func (slice Points) Less(i, j int) bool {
	return slice[i][0] < slice[j][0]
}

func (slice Points) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func FindMinArrowShots(points Points) int {
	// number of non-overlapping intervals
	sort.Sort(points)

	cnt := 1
	end := points[0][1]

	for _, p := range points[1:] {
		// 重疊區間跳過
		if p[0] <= end {
			continue
		}
		cnt++
		end = p[1]
	}

	return cnt
}
