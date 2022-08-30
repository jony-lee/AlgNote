package leetcode

/*
标题 : Combination Sum IV
链接 : https://leetcode.cn/problems/combination-sum-iv/
难度 : 一般
解题思路:
从大到小遍历,如果最大数超过目标,则跳过,否则添加到list中,数据可靠性

注意事项:

*/
type Node struct {
	left, right *Node
	start, end  int
}

type MyCalendar struct {
	root *Node
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (node *Node) insert(start int, end int) bool {
	// 把判断和插入合并到一块了,虽然这个是可以这么做,但是不太建议,扩展性不好.
	if node.start >= end { // 说明在
		if node.left == nil {
			node.left = &Node{start: start, end: end}
			return true
		}
		return node.left.insert(start, end)
	}
	if node.end <= start {
		if node.right == nil {
			node.right = &Node{start: start, end: end}
			return true
		}
		return node.right.insert(start, end)
	}
	return false
}

func (this *MyCalendar) Book(start int, end int) bool {
	if this.root == nil {
		this.root = &Node{start: start, end: end}
		return true
	}
	return this.root.insert(start, end)
}
