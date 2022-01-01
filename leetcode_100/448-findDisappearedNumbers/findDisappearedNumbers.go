package main

import (
	"fmt"
)

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	for _, num := range nums {
		num = (num - 1) % n
		nums[num] += n
	}
	res := make([]int, 0)
	for idx, num := range nums {
		if num <= n {
			res = append(res, idx+1)
		}
	}
	return res
}

func main() {
	//fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDisappearedNumbers([]int{2, 2}))
}
