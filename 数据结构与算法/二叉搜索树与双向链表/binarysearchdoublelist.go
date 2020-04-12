package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var tail *TreeNode = nil

func ConvertRecursion(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 如果左子树为空，则根结点root为双向链表的头结点
	head := ConvertRecursion(root.Left)
	if head == nil {
		head = root
	}

	// 连接当前结点root和当前链表的尾结点tail
	root.Left = tail
	if tail != nil {
		tail.Right = root
	}
	tail = root

	// 遍历当前结点的右子树
	ConvertRecursion(root.Right)

	return head
}

func main() {
	// 	   4
	//   3   5
	// 2
	node2 := &TreeNode{2, nil, nil}
	node3 := &TreeNode{3, node2, nil}
	node5 := &TreeNode{5, nil, nil}
	node4 := &TreeNode{4, node3, node5}
	head := ConvertRecursion(node4)
	for head != nil {
		fmt.Println()
		fmt.Printf("curval: %d", head.Val)
		head = head.Right
	}
	for tail != nil {
		fmt.Println()
		fmt.Printf("curval: %d", tail.Val)
		tail = tail.Left
	}
}
