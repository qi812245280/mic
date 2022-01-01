package main

import "fmt"

func moveZeroes(nums []int) {
	cur := 0
	for idx, num := range nums {
		if num != 0 {
			nums[cur], nums[idx] = nums[idx], nums[cur]
			cur += 1
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
