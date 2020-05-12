package quicksort

import (
	"math/rand"
	"time"
)

func QuickSort(nums []int) {
	// 打乱顺序
	shuffle(nums)

	// 排序
	sort(nums, 0, len(nums)-1)
}

func sort(nums []int, l, h int) {
	if h <= l {
		return
	}

	j := partition(nums, l, h)
	sort(nums, l, j-1)
	sort(nums, j+1, h)
}

func shuffle(nums []int) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
}

func partition(nums []int, l, h int) int {
	i, j := l, h+1
	val := nums[l]

	for {
		for ; nums[i] < val && i != h; i++ {
		}
		for ; nums[j] > val && j != h; j-- {
		}

		if i >= j {
			break
		}

		nums[i], nums[j] = nums[j], nums[i]
	}

	nums[l], nums[j] = nums[j], nums[l]

	return j
}

func QuickSort1(nums []int, l, h int) {
	if l < h {
		i, j := l, h
		pivot := nums[l]

		for i < j {
			for i < j && nums[j] >= pivot {
				j--
			}

			if i < j {
				nums[i] = nums[j]
				i++
			}

			for i < j && nums[i] <= pivot {
				i++
			}

			if i < j {
				nums[j] = nums[i]
				j--
			}
		}

		nums[i] = pivot
		QuickSort1(nums, l, i-1)
		QuickSort1(nums, i+1, h)
	}
}

func QuickSort2(nums []int, low, high int) {
	l := low
	h := high
	pivot := nums[(low+high)/2]
	for l < h {
		for nums[l] < pivot {
			l++
		}
		for nums[h] > pivot {
			h--
		}
		if l >= h {
			break
		}
		nums[l], nums[h] = nums[h], nums[l]
		// 中间有和中位数相等的数
		if nums[l] == pivot {
			l++
		}
		if nums[h] == pivot {
			h--
		}
		if l == h {
			l++
			h--
		}
		if low < h {
			QuickSort2(nums, low, h)
		}
		if high > l {
			QuickSort2(nums, l, high)
		}
	}
}
