package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	//sort.Ints(nums)
	//index := len(nums) / 2
	//return nums[index]

	var count, target int
	for _, num := range nums {
		if count == 0 {
			target = num
		}
		if num == target {
			count++
		} else {
			count--
		}
	}
	return target
}

func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
