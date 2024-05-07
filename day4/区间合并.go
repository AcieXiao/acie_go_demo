package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	start, end int
}

func mergeIntervals(intervals []Interval) []Interval {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})
	var result []Interval
	for _, interval := range intervals {
		if len(result) == 0 || result[len(result)-1].end < interval.start {
			result = append(result, interval)
		} else {
			result[len(result)-1].end = max(result[len(result)-1].end, interval.end)
		}
	}
	return result
}

func main() {
	intervals := []Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}, {9, 17}}
	mergedIntervals := mergeIntervals(intervals)
	for _, interval := range mergedIntervals {
		fmt.Println(interval)
	}
}
