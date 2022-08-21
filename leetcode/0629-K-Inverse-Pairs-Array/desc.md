给出两个整数 `n` 和 `k`，找出所有包含从 `1` 到 `n` 的数字，且恰好拥有 `k` 个逆序对的不同的数组的个数。

逆序对的定义如下：对于数组的第`i`个和第 `j`个元素，如果满`i` < `j`且 `a[i]` > `a[j]`，则其为一个逆序对；否则不是。

由于答案可能很大，只需要返回 答案 mod 109 \+ 7 的值。

**示例 1:**

    
    
    **输入:** n = 3, k = 0
    **输出:** 1
    **解释:** 
    只有数组 [1,2,3] 包含了从1到3的整数并且正好拥有 0 个逆序对。
    

**示例 2:**

    
    
    **输入:** n = 3, k = 1
    **输出:** 2
    **解释:** 
    数组 [1,3,2] 和 [2,1,3] 都有 1 个逆序对。
    

**说明:**

  1.  `n` 的范围是 [1, 1000] 并且 `k` 的范围是 [0, 1000]。
