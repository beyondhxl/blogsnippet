package main

import "fmt"

func InverseParis(nums []int) int {
	// 存储排序好的数组
	tmp := make([]int, len(nums))

	// 合并两个排序好的数组
	var merge func(start, mid, end int) int
	merge = func(start, mid, end int) int { // 匿名函数
		if start >= end {
			return 0
		}
		p1, p2 := mid, end
		k, count := 0, 0
		for p1 >= start && p2 >= mid+1 {
			// 第一个数组中的数字不大于第二个数组中的数字，则进行下一轮的比较
			if nums[p1] <= nums[p2] {
				tmp[k] = nums[p2]
				p2--
				k++
			} else {
				// 两个数组合并时已排好序，对于num[p1]而言，至少有(p2-mid)个逆序对
				tmp[k] = nums[p1]
				count += p2 - mid
				p1--
				k++
			}
		}
		// p1指向的数字此时都不大于p2指向的数字
		for p1 >= start {
			tmp[k] = nums[p1]
			p1--
			k++
		}
		for p2 >= mid+1 {
			tmp[k] = nums[p2]
			p2--
			k++
		}
		for i := 0; i <= k-1; i++ {
			nums[end-i] = tmp[i]
		}

		return count
	}

	var sort func(start, end int) int
	sort = func(start, end int) int {
		count := 0
		if start < end {
			// 每次排序都从中间分割
			mid := (start + end) / 2
			count += sort(start, mid)
			count += sort(mid+1, end)
			count += merge(start, mid, end)
		}
		return count
	}
	return sort(0, len(nums)-1)
}

func main() {
	test := []int{7, 5, 6, 4}
	pairs := InverseParis(test)
	fmt.Println("样例输入：")
	for i := 0; i < len(test); i++ {
		fmt.Println(test[i])
	}
	fmt.Printf("逆序对个数：%d", pairs)
}
