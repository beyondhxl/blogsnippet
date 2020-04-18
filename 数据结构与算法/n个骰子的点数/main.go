package main

import (
	"fmt"
	"math"
)

/*
func printProbability(n int) []string {
	if n <= 0 {
		return nil
	}

	total := math.Pow(6, float64(n))
	var result []string

	// dp[c][k] c个骰子朝上一面点数之和为k的次数
	var dp [][]int

	// 初始化dp[1][1...6]
	for i := 1; i <= 6; i++ {
		dp[1][i] = 1
	}

	sum := 0
	for i := 2; i <= n; i++ {
		for j := 2; j <= 6*n; j++ {
			for m := 1; m < j && m <= 6; m++ {
				sum += dp[i-1][j-m]
			}
			dp[i][j] = sum
		}
	}
	var numstr string
	for k := n; k <= 6*n; k++ {
		numstr = strconv.Itoa(dp[n][k])
		result = append(result, numstr + string('/') + strconv.Itoa(int(total)))
	}
	return result
}
*/

// 动态规划解法
const gmax = 6

func printProbability(n int) (result []float64) {
	var (
		i, j, k       int
		total         = math.Pow(float64(gmax), float64(n))
		flag          = 0
		probabilities = make([][]int, 2)
	)
	probabilities[0] = make([]int, n*gmax+1)
	probabilities[1] = make([]int, n*gmax+1)

	for i = 1; i <= gmax; i++ {
		probabilities[flag][i] = 1
	}

	for k = 2; k <= n; k++ {
		for i = 0; i < k; i++ {
			probabilities[1-flag][i] = 0
		}

		for i = k; i <= k*gmax; i++ { // 使用i个骰子最小点数为i，最大点数6*i
			probabilities[1-flag][i] = 0
			for j = 1; i-j >= 0 && j <= gmax; j++ { // 第j个骰子的6种情况
				probabilities[1-flag][i] += probabilities[flag][i-j]
			}
		}
		flag = 1 - flag
	}

	result = make([]float64, gmax*n-n+1)
	for i = n; i <= gmax*n; i++ {
		result[i-n] = float64(probabilities[flag][i]) / total
	}

	return result
}

func main() {
	result := printProbability(3)
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}
