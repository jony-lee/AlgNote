package main

// @File    :   main.go
// @Time    :   2021/02/15 23:26:05
// @Author  :   jony
// @Contact :   jonylee_cn@qq.com

//https://leetcode-cn.com/problems/replace-words/
/*
该提分两步,
*/
// import "strings"
import (
	"fmt"
	"sort"
	"strings"
)
// 最佳解法
func replaceWords(dictionary []string, sentence string) string {
	// 对字符串进行排序
	sort.Strings(dictionary)
	arr := strings.Split(sentence, " ")
	for i:= 0;i<len(arr);i++{
		for _,v := range dictionary{
			if strings.HasPrefix(arr[i],v){
				arr[i] = v
				break
			}
		}
		
	}
	return strings.Join(arr," ")
}

// 最初版本
func replaceWords_old1(dictionary []string, sentence string) string {
	// 对字符串进行排序
	sort.Strings(dictionary)
	str := []string{}
	for _, v := range strings.Split(sentence, " ") {
		str = append(str, replaceWord(dictionary, v))
	}
	return strings.Join(str, " ")
}

// old 替换一个单词
func replaceWord(dictionary []string, word string) string {
	for i := 0; i < len(dictionary); i++ {
		if strings.HasPrefix(word,dictionary[i]) {
			return dictionary[i]
		}
	}
	return word
}

// main ...
func main() {
	x := []string{"a", "b", "c"}
	a := replaceWords(x, "aadsfasf absbs bbab cadsfafs")
	fmt.Println(a)
}
