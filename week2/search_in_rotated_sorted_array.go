// 整数数组 nums 按升序排列，数组中的值 互不相同 。

// 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。

// 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。

package main

import "fmt"

func main() {
    fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0)) // 4
    fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3)) // -1
}

// 类二分搜索
// 最左边数 < 中间数则左侧有序，最右边数 > 中间数则右侧有序
// 在缩小搜索区域时，一直只在确定的有序区域内查找
func search(nums []int, target int) int {
    l, r := 0, len(nums)-1
    for l <= r {
        mid := (l + r) / 2
        switch {
        case nums[mid] == target: // bingo
            return mid
        case nums[l] <= nums[mid]: // 左侧有序
            if nums[l] <= target && target < nums[mid] { // 保证 target 一定在有序的左侧内
                r = mid - 1
            } else {
                l = mid + 1
            }
        case nums[mid] <= nums[r]: // 右侧有序
            if nums[mid] < target && target <= nums[r] { // 保证 target 一定在有序的右侧内
                l = mid + 1
            } else {
                r = mid - 1
            }
        }
    }
    return -1
}
