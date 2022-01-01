package main

import "fmt"

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	var dfs func(int, []int)
	dfs = func(start int, path []int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(i+1, path)
			path = path[:len(path)-1]
		}
	}
	dfs(0, []int{})
	return res
}

func main() {
	path := []int{1, 2, 3}
	fmt.Println(subsets(path))
}
