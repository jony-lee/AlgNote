package leetcode

// @File    :   main.go
// @Time    :   2021/02/15 17:34:43
// @Author  :   jony
// @Contact :   jonylee_cn@qq.com

import "sort"

// 最佳解法
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	ret := [][]int{}
	//sl是变化的list
	sl := make([]int, len(nums))
	dfs(nums, &ret, sl, 0, 0)
	return ret
}

// start起始位置
func dfs(nums []int, ret *[][]int, sl []int, start int, layer int) {
	dst := make([]int, layer)
	copy(dst, sl)
	*ret = append(*ret, dst)
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		sl[layer] = nums[i]
		dfs(nums, ret, sl, i+1, layer+1)
	}
}
