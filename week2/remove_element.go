// 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。

// 不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。

// 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

// 说明:

// 为什么返回数值是整数，但输出的答案是数组呢?

// 请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。

package main

import "fmt"

func main() {
    fmt.Println(removeElement1([]int{3, 2, 2, 3}, 3))
    fmt.Println(removeElement2([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

//
// 遍历数组一如既往地联想到双指针法
//
func removeElement1(nums []int, val int) int {
    slow := 0
    for fast := 0; fast < len(nums); fast++ {
        if nums[fast] != val {
            nums[slow] = nums[fast] // 快指针在前，慢指针在后，遇到不等的元素就放到慢指针的位置
            slow++
        }
    }
    return slow
}

// 类似 26 的解法
func removeElement2(nums []int, val int) int {
    slow, fast := 0, 0
    for fast < len(nums) {
        if nums[fast] != val {
            if slow != fast {
                nums[slow] = nums[fast]
            }
            slow++
            fast++
            continue
        }
        fast++
    }
    return slow
}
