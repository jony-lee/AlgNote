**n 位格雷码序列** 是一个由 `2n` 个整数组成的序列，其中：

  * 每个整数都在范围 `[0, 2n - 1]` 内（含 `0` 和 `2n - 1`）
  * 第一个整数是 `0`
  * 一个整数在序列中出现 **不超过一次**
  * 每对 **相邻** 整数的二进制表示 **恰好一位不同** ，且
  * **第一个** 和 **最后一个** 整数的二进制表示 **恰好一位不同**

给你一个整数 `n` ，返回任一有效的 **n 位格雷码序列** 。



**示例 1：**

    
    
    **输入：** n = 2
    **输出：** [0,1,3,2]
    **解释：**
    [0,1,3,2] 的二进制表示是 [00,01,11,10] 。
    - 0 ** _0_** 和 0 _ **1**_ 有一位不同
    - _**0**_ 1 和 _**1**_ 1 有一位不同
    - 1 _ **1**_ 和 1 _ **0**_ 有一位不同
    - _**1**_ 0 和 _**0**_ 0 有一位不同
    [0,2,3,1] 也是一个有效的格雷码序列，其二进制表示是 [00,10,11,01] 。
    - _**0**_ 0 和 _**1**_ 0 有一位不同
    - 1 _ **0**_ 和 1 _ **1**_ 有一位不同
    - _**1**_ 1 和 _**0**_ 1 有一位不同
    - 0 _ **1**_ 和 0 _ **0**_ 有一位不同
    

**示例 2：**

    
    
    **输入：** n = 1
    **输出：** [0,1]
    



**提示：**

  * `1 <= n <= 16`
