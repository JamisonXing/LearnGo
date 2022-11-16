package day6

import "container/list"

//https://www.bilibili.com/video/BV1U34y1D7MN/?spm_id_from=333.999.0.0&vd_source=1761deea5b8c69ccd1cd089a400b2c54

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		tmp := queue.Front().Value.(*TreeNode)
		ans = append(ans, tmp.Val)
		queue.Remove(queue.Front())
		if tmp.Left != nil {
			queue.PushBack(tmp.Left)
		}
		if tmp.Right != nil {
			queue.PushBack(tmp.Right)
		}
	}
	return ans
}

/*
func levelOrder(root *TreeNode) []int {
    var ans []int
    if root == nil {
        return ans
    }

    queue := list.New()
    queue.PushBack(root)

    for queue.Len() != 0 {
        temp := queue.Front().Value.(*TreeNode)
        ans = append(ans, temp.Val)
        queue.Remove(queue.Front())
        if temp != nil {
            queue.PushBack(temp.Left)
        }

        if temp != nil {
            queue.PushBack(temp.Right)
        }
    }

    return ans

}*/
