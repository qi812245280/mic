package main

import "fmt"

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	var dfs func(int, []int)
	dfs = func(start int, path []int) {
		if len(path) == k {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := start; i <= n; i++ {
			path = append(path, i)
			dfs(i+1, path)
			path = path[:len(path)-1]
		}
	}
	dfs(1, []int{})
	return res
}

func main() {
	fmt.Println(combine(4, 2))
}
