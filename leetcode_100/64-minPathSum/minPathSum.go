package main

import (
	"fmt"
	"math"
)

func minPathSum(grid [][]int) int {
	res := math.MaxInt64
	ilength := len(grid)
	jlength := len(grid[0])
	var dfs func(int, int, int)
	dfs = func(i, j, path int) {
		if i == ilength-1 && j == jlength-1 {
			path += grid[i][j]
			res = int(math.Min(float64(res), float64(path)))
			return
		}
		path += grid[i][j]
		if i < ilength-1 {
			dfs(i+1, j, path)
		}
		if j < jlength-1 {
			dfs(i, j+1, path)
		}
	}
	dfs(0, 0, 0)
	return res
}
func minPathSum2(grid [][]int) int {
	ilength := len(grid)
	jlength := len(grid[0])
	for i := 0; i < ilength; i++ {
		for j := 0; j < jlength; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 && j > 0 {
				grid[i][j] += grid[i][j-1]
			}
			if i > 0 && j == 0 {
				grid[i][j] += grid[i-1][j]
			}
			if i > 0 && j > 0 {
				grid[i][j] = grid[i][j] + int(math.Min(float64(grid[i][j-1]), float64(grid[i-1][j])))
			}
		}
	}
	return grid[ilength-1][jlength-1]
}

func main() {
	fmt.Println(minPathSum2([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
	fmt.Println(minPathSum2([][]int{{1, 2, 3}, {4, 5, 6}}))
}
