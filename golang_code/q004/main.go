package main

import (
	"fmt"
)

/*
判断两个给定的字符串排序后是否一致
问题描述

给定两个字符串，请编写程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。 这里规定【大小写为不同字符】，且考虑字符串重点空格。给定一个string s1和一个string s2，请返回一个bool，代表两串是否重新排列后可相同。 保证两串的长度都小于等于5000。

解题思路

首先要保证字符串长度小于5000。之后只需要一次循环遍历s1中的字符在s2是否都存在即可。
*/

func isRegroup(s1, s2 string) bool {
	//for _, le := range s1 {
	//	if strings.Count(s1, string(le)) != strings.Count(s2, string(le)) {
	//		return false
	//	}
	//}
	//return true
	if len(s1) != len(s2) || len(s1) == 0 {
		return false
	}
	var res = s1[0] ^ s2[0]
	for i := 1; i < len(s1); i++ {
		res = res ^ s1[i] ^ s2[i]
	}
	if res != 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isRegroup("哥哥爱妹妹", "爱妹哥哥妹"))
}
