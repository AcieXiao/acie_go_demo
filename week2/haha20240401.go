package main

func Find(_list []int, target int) int {
	//for k, v := range _list {
	//	if v==target{
	//		return k
	//	}
	//}
	//return -1
	left, right := 0, len(_list)

	for left <= right {
		midIndex := left + (right-left)/2
		if _list[midIndex] == target {
			return midIndex
		} else if _list[midIndex] < target {
			left = midIndex + 1
		} else if _list[midIndex] > target {
			right = midIndex - 1
		}
	}
	return -1
}
