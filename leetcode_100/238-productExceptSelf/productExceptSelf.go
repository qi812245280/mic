package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	res[0] = 1
	for i := 1;i <length;i++{
		res[i] = res[i-1]*nums[i-1]
	}
	tmp2 := 1
	for i := length -1; i>=0; i--{
		res[i] = tmp2 * res[i]
		tmp2 = tmp2 *nums[i]
	}
	return res
}

func main()  {
	fmt.Println(productExceptSelf([]int{1,2,3,4}))
}
