package main

import "fmt"

// 断言函数指针
type Pred func(int) bool

// 是否是偶数
func IsEven(n int) bool {
	return n&1 == 0
}

// 奇数在前断言函数
func oddFirstPred(s []int, pred Pred) {
	left, right := 0, len(s)-1
	for left < right {
		for pred(s[right]) == false && left < right {
			right--
		}
		for pred(s[left]) == true && left < right {
			left++
		}
		if left == right {
			break
		}
		if left < right {
			s[left], s[right] = s[right], s[left]
		}
	}
}

// 奇数在前
func OddFirst(s []int) {
	oddFirstPred(s, IsEven)
}

func main() {
	l := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("before: ", l)
	OddFirst(l)
	fmt.Println("after: ", l)
}
