package main

import (
	"log"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

type IntervalSlice []Interval

func (intervals IntervalSlice) Len() int {
	return len(intervals)
}

func (intervals IntervalSlice) Less(i, j int) bool {
	return intervals[i].End < intervals[j].End
}

func (intervals IntervalSlice) Swap(i, j int) {
	intervals[i], intervals[j] = intervals[j], intervals[i]
}

func eraseOverlapIntervals(intervals IntervalSlice) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Sort(intervals)

	// number of non-overlapping intervals
	end := intervals[0].End
	cnt := 1

	for _, interval := range intervals {
		if interval.Start < end {
			continue
		}

		end = interval.End
		cnt++
	}

	// number of overlapping intervals
	return len(intervals) - cnt
}

func main() {
	intervals := IntervalSlice{
		Interval{
			Start: 1,
			End:   100,
		},
		Interval{
			Start: 2,
			End:   50,
		},
		Interval{
			Start: 1,
			End:   10,
		},
		Interval{
			Start: 1,
			End:   2,
		},
	}

	log.Println(intervals)
}
