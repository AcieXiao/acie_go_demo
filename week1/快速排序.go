package main

import "fmt"

//func partition(list []int, low, height int) int {
//	pivot := list[height]
//	_low, _height := low, height
//	i := 0
//	for low < height {
//		i++
//		for low < height && pivot >= list[low] {
//			low++
//		}
//		list[height] = list[low]
//		fmt.Println(_low, " ", _height, " 第", i, "次排序 111", list)
//
//		for low < height && pivot <= list[height] {
//			height--
//		}
//		list[low] = list[height]
//		fmt.Println(_low, " ", _height, " 第", i, "次排序 222", list)
//	}
//	list[height] = pivot
//	fmt.Println(_low, " ", _height, " 第", i, "次排序 333", list)
//	return height
//}

func partition(list []int, low, height int) int {
	povit := list[height]
	for low < height {
		for low < height && povit >= list[low] {
			low++
		}
		list[height] = list[low]

		for low < height && povit <= list[height] {
			height--
		}
		list[low] = list[height]
	}
	list[height] = povit
	return height
}

func quickSort(list []int, low, height int) {
	if height > low {
		pivot := partition(list, low, height)
		quickSort(list, low, pivot-1)
		quickSort(list, pivot+1, height)
	}
}

func main() {
	//list := []int{2, 44, 4, 8, 33, 1, 22, -11, 6, 34, 7, 54, 9}
	list := []int{211, 44, 50, 50}
	quickSort(list, 0, len(list)-1)
	fmt.Println(list)
}
