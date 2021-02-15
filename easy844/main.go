package main

// @File    :   main.go
// @Time    :   2021/02/16 00:14:36
// @Author  :   jony
// @Contact :   jonylee_cn@qq.com

func backspaceCompare(S string, T string) bool {
	return finalString(S) == finalString(T)
}

func finalString(s string) string {
	ret := []rune{}
	for _, v := range s {
		if v != '#' {
			ret = append(ret, v)
		} else if len(ret) > 0 {
			ret = ret[:len(ret)-1]
		}
	}
	return string(ret)
}
