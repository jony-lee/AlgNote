package leetcode

import (
	"strconv"
	"strings"
)

/*
标题 : Compress String LCCI
链接 : https://leetcode.cn/problems/compress-string-lcci/
难度 : 简单
解题思路:
该题比较简单,主要是遍历字符串,建立状态机,
需要注意对边界条件的处理,一个是短字符串的处理,一个是末位字符串的处理.
golang原始的方式可以将string当做[]byte来处理,此处可以通过strings包来简化字符串的添加.

注意事项:

*/

func compressString(S string) string {
	if len(S) <= 2 {
		return S
	}
	buf := strings.Builder{}
	buf.WriteByte(S[0])
	cnt := 1
	for i := 1; i < len(S); i++ {
		if S[i] == S[i-1] {
			cnt++
			continue
		}
		buf.WriteString(strconv.Itoa(cnt))
		buf.WriteByte(S[i])
		cnt = 1
	}
	buf.WriteString(strconv.Itoa(cnt))
	if buf.Len() >= len(S) {
		return S
	}
	return buf.String()
}
