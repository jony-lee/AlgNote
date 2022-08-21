给你一个下标从 **0** 开始的字符串 `s` 。另给你一个下标从 **0** 开始、长度为 `k` 的字符串 `queryCharacters`
，一个下标从 `0` 开始、长度也是 `k` 的整数 **下标** 数组 `queryIndices` ，这两个都用来描述 `k` 个查询。

第 `i` 个查询会将 `s` 中位于下标 `queryIndices[i]` 的字符更新为 `queryCharacters[i]` 。

返回一个长度为 `k` 的数组 `lengths` ，其中 `lengths[i]` 是在执行第 `i` 个查询 **之后** `s` 中仅由
**单个字符重复** 组成的 **最长子字符串** 的 **长度** _。_



**示例 1：**

    
    
    **输入：** s = "babacc", queryCharacters = "bcb", queryIndices = [1,3,3]
    **输出：** [3,3,4]
    **解释：**
    - 第 1 次查询更新后 s = " _b **b** b_acc" 。由单个字符重复组成的最长子字符串是 "bbb" ，长度为 3 。
    - 第 2 次查询更新后 s = "bbb _ **c** cc_" 。由单个字符重复组成的最长子字符串是 "bbb" 或 "ccc"，长度为 3 。
    - 第 3 次查询更新后 s = " _bbb **b**_ cc" 。由单个字符重复组成的最长子字符串是 "bbbb" ，长度为 4 。
    因此，返回 [3,3,4] 。

**示例 2：**

    
    
    **输入：** s = "abyzz", queryCharacters = "aa", queryIndices = [2,1]
    **输出：** [2,3]
    **解释：**
    - 第 1 次查询更新后 s = "ab **a** _zz_ " 。由单个字符重复组成的最长子字符串是 "zz" ，长度为 2 。
    - 第 2 次查询更新后 s = " _a **a** a_zz" 。由单个字符重复组成的最长子字符串是 "aaa" ，长度为 3 。
    因此，返回 [2,3] 。
    



**提示：**

  * `1 <= s.length <= 105`
  * `s` 由小写英文字母组成
  * `k == queryCharacters.length == queryIndices.length`
  * `1 <= k <= 105`
  * `queryCharacters` 由小写英文字母组成
  * `0 <= queryIndices[i] < s.length`

