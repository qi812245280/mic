package main

import (
	"fmt"
)

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := [][]int{}
	for i, temperature := range temperatures{
		if len(stack) == 0 {
			stack = append(stack, []int{i, temperature})
			continue
		}
		for len(stack) >0 {
			top := stack[len(stack) -1]
			if temperature > top[1]{
				res[top[0]] = i - top[0]
				stack = stack[:len(stack) -1]
			} else{
				break
			}
		}
		stack = append(stack, []int{i, temperature})
	}
	for _, nums := range stack{
		res[nums[0]] = 0
	}
	return res
}

func main()  {
	fmt.Println(dailyTemperatures([]int{73,74,75,71,69,72,76,73}))
	fmt.Println(dailyTemperatures([]int{30,40,50,60}))
	fmt.Println(dailyTemperatures([]int{30,60,90}))
}
