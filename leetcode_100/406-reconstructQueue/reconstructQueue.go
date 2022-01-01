package main

import (
	"fmt"
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	// 排序
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0]{
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	res := make([][]int, 0)
	for _, person := range people{
		idx := person[1]
		res = append(res, person)
		copy(res[idx+1:], res[idx:])
		res[idx] = person
	}
	return res
}

func main() {
	fmt.Println(reconstructQueue([][]int{{7,0},{4,4},{7,1},{5,0},{6,1},{5,2}}))
}
