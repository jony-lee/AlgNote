package leetcode

/*
标题 : Maximum Level Sum of a Binary Tree
链接 : https://leetcode.cn/problems/maximum-level-sum-of-a-binary-tree/
难度 : 一般
解题思路:
思路就是使用递归的方法,递归的时候需要计算每一层的加和,因此需要一个数组存储值,然后再比对.


注意事项:
这道题无需考虑前序,中序还是后续遍历.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	sum := make([]int, 100, 100)
	max_level := travel(root, 0, sum)
	res := 0
	max := root.Val
	for i := 0; i <= max_level; i++ {
		if max < sum[i] {
			max = sum[i]
			res = i
		}
	}
	return res + 1
}

func travel(node *TreeNode, level int, sum []int) int {
	if node == nil {
		return level - 1
	}
	sum[level] += node.Val
	l1 := travel(node.Left, level+1, sum)
	l2 := travel(node.Right, level+1, sum)
	if l1 > l2 {
		return l1
	}
	return l2
}
