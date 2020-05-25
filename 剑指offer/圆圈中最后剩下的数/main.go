package main

import "fmt"

// n个数字从0、1、2、...、n-1开始编号，每次从其中删除第m个数字
func CycleNumber(n, m int) int {
	if n < 1 || m < 1 {
		return -1
	} else if n == 1 {
		return 0
	} else {
		return (CycleNumber(n-1, m) + m) % n
	}
}

func main() {
	last := CycleNumber(5, 3)
	fmt.Printf("5个数字编号0、1、2、3、4，每次删除第3个数字，剩下的数字：%d", last)
}
