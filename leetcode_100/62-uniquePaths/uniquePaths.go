package main

import "fmt"

func uniquePaths(m int, n int) int {
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		res[i][0] = 1
	}
	for j := 0; j < n; j++ {
		res[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			res[i][j] = res[i-1][j] + res[i][j-1]
		}
	}
	return res[m-1][n-1]
}

func main() {
	fmt.Println(uniquePaths(3, 7))
}
