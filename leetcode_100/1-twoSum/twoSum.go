package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	mapInfo := make(map[int]int)
	res := make([]int, 0, 2)
	for idx, num := range nums {
		gap := target - num
		if idx2, ok := mapInfo[gap]; ok {
			res = append(res, idx2)
			res = append(res, idx)
			break
		}
		mapInfo[num] = idx
	}
	return res
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Println(twoSum(nums, target))
}
