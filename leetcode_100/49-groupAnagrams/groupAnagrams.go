package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	mapInfo := make(map[string][]string)
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		ss := string(s)
		if _, ok := mapInfo[ss]; ok {
			mapInfo[ss] = append(mapInfo[ss], str)
		} else {
			mapInfo[ss] = []string{str}
		}
	}
	res := make([][]string, 0)
	for _, v := range mapInfo {
		res = append(res, v)
	}
	return res
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
