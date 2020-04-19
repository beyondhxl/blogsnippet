package main

import (
	"fmt"
	"math"
)

// 利用公式1+2+...+n = (n+1)*n/2
func sumSolution(n int) int {
	return (int(math.Pow(float64(n), 2)) + n) >> 1
}

func main() {
	n := 3
	res := sumSolution(n)
	fmt.Printf("当n=%d，1+2+...+n的值是%d", n, res)
}
