package main
// @File    :   main.go
// @Time    :   2021/02/15 15:36:03
// @Author  :   jony
// @Contact :   jonylee_cn@qq.com

//https://leetcode-cn.com/problems/max-consecutive-ones/submissions/

func findMaxConsecutiveOnes(nums []int) int {
    var max = 0
    var length = 0
    for i:= 0;i< len(nums);i++{
        if nums[i]==1{
            length++
        }
        if max < length {
            max = length
        }
        if nums[i]!= 1{
            length = 0
        }
    }
    return max
}
