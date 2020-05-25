package main

import "fmt"

func oneNumber(nums []int) []int {
	xor := 0
	// 遍历数组进行异或操作
	for _, v := range nums {
		xor ^= v
	}

	// 取xor从右往左第一位为1的位
	// 与自身的负数相与
	// 负数的补码：符号位为1，其余位为该数绝对值的原码按位取反，然后再加1
	// 2=0010 -2=1101+1=1110 0010&1110=0010
	// 用来判断xor是不是2的N次方
	lowbit := xor & -xor

	one, two := 0, 0
	for _, v := range nums {
		if lowbit&v == 0 {
			one ^= v
		} else {
			two ^= v
		}
	}
	return []int{one, two}
}

func main() {
	test := []int{2, 4, 3, 6, 3, 2, 5, 5}
	result := oneNumber(test)
	fmt.Print(result)
}
