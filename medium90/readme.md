# 题目
https://leetcode-cn.com/problems/subsets-ii/

# 题目分析
该题是一个遍历题

# 解题思路

遍历题可以用深度优先遍历来解决
dfs的参数包括

nums是原始数组,ret是待返回数组,sl是深度遍历时的附加元素
当完成整个dfs之后,所有的数组通过append的方式加入到了ret里面

func dfs(nums []int, ret *[][]int, sl []int, start int, layer int) {
	dst := make([]int, layer)
	copy(dst, sl)
	*ret = append(*ret, dst)
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue //跳过元素相同的值
		}
		sl[layer] = nums[i]
		dfs(nums, ret, sl, i+1, layer+1)
	}
}