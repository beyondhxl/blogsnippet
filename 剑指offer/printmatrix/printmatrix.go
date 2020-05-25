package main

import "fmt"

func PrintMatrix(matrix [][]int) []int {
	var res []int
	// 行数
	n := len(matrix)
	if n == 0 {
		return res
	}
	// 列数
	m := len(matrix[0])
	// 计算元素个数
	nm := n * m
	// 初始化层
	layer := 0
	// 初始化边界
	startN, endN, startM, endM := 0, 0, 0, 0
	// 从外围到内部顺时针遍历
	for nm > 0 {
		startN, endN = layer, n-layer-1
		startM, endM = layer, m-layer-1

		// 从左到右
		for i := startM; i <= endM && nm > 0; i++ {
			res = append(res, matrix[startN][i])
			nm--
		}

		// 从上到下
		for i := startN + 1; i <= endN && nm > 0; i++ {
			res = append(res, matrix[i][endM])
			nm--
		}

		// 从右到左
		for i := endM - 1; i >= startM && nm > 0; i-- {
			res = append(res, matrix[endN][i])
			nm--
		}

		// 从下到上
		for i := endN - 1; i >= startN+1 && nm > 0; i-- {
			res = append(res, matrix[i][startM])
			nm--
		}

		// 层数递增
		layer++
	}
	return res
}

func main() {
	var matrix [][]int = [][]int{{1, 2}, {4, 5}, {7, 8}}
	res := PrintMatrix(matrix)
	fmt.Print(res)
}
