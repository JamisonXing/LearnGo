package main

import "fmt"

//https://leetcode.cn/problems/symmetric-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 官方解答 递归
func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(q.Left, p.Right)
}

func main() {
	//root = [1,3,2],不对称
	rightRoot := TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	leftRoot := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	root := TreeNode{
		Val:   1,
		Left:  &leftRoot,
		Right: &rightRoot,
	}

	if isSymmetric(&root) {
		fmt.Println("对称")
	} else {
		fmt.Println("不对称")
	}

}
