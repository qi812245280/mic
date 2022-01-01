package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	var dfs func(int, int, []int)
	dfs = func(start, gap int, path []int) {
		if gap <= 0 {
			if gap == 0 {
				tmp := make([]int, len(path))
				copy(tmp, path)
				res = append(res, tmp)
			}
			return
		}
		for i := start; i < len(candidates); i++ {
			path = append(path, candidates[i])
			dfs(i, gap-candidates[i], path)
			path = path[:len(path)-1]
		}
	}
	dfs(0, target, []int{})
	return res
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}
