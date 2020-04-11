package main

import (
	"fmt"
)

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

func FindPath(root *TreeNode, sum int) {
	if root == nil || root.Data == 0 || root.Left == nil || root.Right == nil {
		return
	}

	var path []int
	cursum := 0
	findPath(root, sum, path, cursum)
	fmt.Println("------")
	fmt.Println(path)
	fmt.Println(cursum)
}

func findPath(root *TreeNode, sum int, path []int, cursum int) {
	cursum += root.Data
	path = append(path, root.Data)

	// 如果是叶子节点，且路径上节点的和等于输入的值
	// 打印出这条路径
	leaf := root.Left == nil && root.Right == nil
	if cursum == sum && leaf {
		fmt.Print("a path is found: ")
		for i := 0; i < len(path); i++ {
			fmt.Printf("%d --> ", path[i])
		}
		fmt.Println()
	}

	// 如果不是叶子节点，则遍历它的子节点
	if root.Left != nil {
		findPath(root.Left, sum, path, cursum)
	}
	if root.Right != nil {
		findPath(root.Right, sum, path, cursum)
	}

	// 在返回到父节点之前，在路径上删除当前节点，
	// 并在cursum中减去当前节点的值
	//cursum -= root.Data
	//path = path[:len(path)-1]
	fmt.Println(cursum)
	fmt.Println(path)
}

func main() {
	node4 := &TreeNode{4, nil, nil}
	node7 := &TreeNode{7, nil, nil}
	node5 := &TreeNode{5, node4, node7}
	node12 := &TreeNode{12, nil, nil}
	node10 := &TreeNode{10, node5, node12}
	FindPath(node10, 22)
}
