package main

import "fmt"

func numberCount(nums []int, target int) int {
	// 二分查找初始化游标
	left, right := 0, len(nums)-1
	var mid int
	for left < right {
		mid = (left + right) / 2
		// 中间元素与查找对象相等，继续在它两边查找
		if nums[mid] == target {
			// right在末尾，与中间元素不等，则往前面查找
			for nums[mid] != nums[right] {
				right--
			}
			// left在头部，与中间元素不等，则往后顺序查找
			for nums[mid] != nums[left] {
				left++
			}
			break
		}
		if nums[mid] > target {
			// 中间元素比目标值大，则在数组的前半段去查找
			right = mid - 1
		} else {
			// 中间元素比目标值大，则在数组的后半段去查找
			left = mid + 1
		}
	}
	if left < right {
		return right - left + 1
	}
	return -1
}

func main() {
	test := []int{1, 2, 3, 3, 3, 3, 4, 5}
	count := numberCount(test, 3)
	fmt.Println("Count 3 in ", test, ": ", count)
}
