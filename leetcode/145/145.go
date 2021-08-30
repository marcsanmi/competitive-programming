package main

import "fmt"

// Binary Tree Postorder Traversal

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversalIt(root *TreeNode) []int {
	output := make([]int, 0, 0)
	stack := make([]*TreeNode, 0, 0)

	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			output = append([]int{root.Val}, output...)
			root = root.Right
		} else {
			root = stack[len(stack)-1].Left
			stack = stack[:len(stack)-1]
		}
	}

	return output
}

func postorderTraversalRec(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	if root.Left != nil && root.Right != nil {
		return append(append(postorderTraversalRec(root.Left), postorderTraversalRec(root.Right)...), root.Val)
	}

	if root.Left != nil {
		return append(postorderTraversalRec(root.Left), root.Val)
	}

	return append(postorderTraversalRec(root.Right), root.Val)
}

func main() {
	treeNode := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 4,
			},
			Left: &TreeNode{Val: 7},
		},
	}

	fmt.Println(postorderTraversalIt(&treeNode))
}
