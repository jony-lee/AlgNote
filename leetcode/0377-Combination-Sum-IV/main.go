package leetcode

/*
标题 : Combination Sum IV
链接 : https://leetcode.cn/problems/combination-sum-iv/
难度 : 一般
解题思路:
从大到小遍历,如果最大数超过目标,则跳过,否则添加到list中,数据可靠性

注意事项:

*/
type Set map[int]struct{}

func (s Set) Add(ele int) {
	s[ele] = struct{}{}
}

func (s Set) Has(ele int) bool {
	_, ok := s[ele]
	return ok
}

type MyCalendar struct {
	// 此处利用了动态线段树的完全二叉树的特性,通过一个map来动态存储线段树的节点
	// 否则其实使用静态线段树在建树的时候就已经确定了数组长度,此时应当使用数组
	// 此处其实只需要集合类型,来判断当前元素有还是没有,因此可以构造一个Set
	tree, lazy Set
}

func Constructor() MyCalendar {
	return MyCalendar{Set{}, Set{}}
}

func (c MyCalendar) query(start, end, l, r, idx int) bool {
	if r < start || end < l {
		return false
	}
	if c.lazy.Has(idx) { // 如果该区间已被预订，则直接返回
		return true
	}
	if start <= l && r <= end {
		return c.tree.Has(idx)
	}
	mid := (l + r) >> 1
	return c.query(start, end, l, mid, 2*idx) ||
		c.query(start, end, mid+1, r, 2*idx+1)
}

func (c MyCalendar) update(start, end, l, r, idx int) {
	if r < start || end < l { // 请求区间段在当前线段树之外
		return
	}
	if start <= l && r <= end { // 请求区间覆盖线段树区间
		c.tree.Add(idx)
		c.lazy.Add(idx)
	} else {
		mid := (l + r) >> 1
		c.update(start, end, l, mid, 2*idx)     // 更新完全二叉树的左子树
		c.update(start, end, mid+1, r, 2*idx+1) // 更新完全二叉树的右子树
		c.tree.Add(idx)
		if c.lazy.Has(2*idx) && c.lazy.Has(2*idx+1) {
			c.lazy.Add(idx) // 如果有子节点的话,就懒标记为当前子树被预定了.
		}
	}
}

func (c MyCalendar) Book(start, end int) bool {
	if c.query(start, end-1, 0, 1e9, 1) {
		return false
	}
	c.update(start, end-1, 0, 1e9, 1)
	return true
}
