package day6

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder2(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	/*nodeList:=[] *TreeNode{root} //节点数组用于节点的出入栈
	  for len(nodeList)!=0{
	      l:=len(nodeList) // 长度为多少即该层有多少元素
	      levellist:=[]int{}
	      for i:=0;i<l;i++{
	          node:=nodeList[0]
	          levellist=append(levellist,node.Val) //值加入levellist
	          nodeList=nodeList[1:]  //删除
	          if node.Left != nil{  //左右节点加入nodelist
	                nodeList = append(nodeList, node.Left)
	            }
	          if node.Right != nil{
	                nodeList = append(nodeList, node.Right)
	            }
	     }
	      res=append(res,levellist)  //该层的值数组加入res*/

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() != 0 {
		size := queue.Len()
		var temp []int

		for i := 0; i < size; i++ {
			node := queue.Front().Value.(*TreeNode)
			temp = append(temp, node.Val)
			queue.Remove(queue.Front())

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		res = append(res, temp)
	}

	return res
}
