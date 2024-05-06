package main

import "fmt"

func maopaopaixu(arr []int) []int {
	n := len(arr)

	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

func main() {
	oriList := []int{2, 1, 3, 5, 4, 9, 6, 7, 8}
	fmt.Println(maopaopaixu(oriList))
}
