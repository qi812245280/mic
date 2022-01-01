package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	var dfs func(int, int, []int)
	dfs = func(start, gap int, path []int) {
		if gap <= 0 {
			if gap == 0 && len(path) == k {
				tmp := make([]int, k)
				copy(tmp, path)
				res = append(res, tmp)
			}
			return
		}
		for i := start; i <= 9; i++ {
			path = append(path, i)
			dfs(i+1, gap-i, path)
			path = path[:len(path)-1]
		}
	}
	dfs(1, n, []int{})
	return res
}

func main() {
	fmt.Println(combinationSum3(3, 9))
}
