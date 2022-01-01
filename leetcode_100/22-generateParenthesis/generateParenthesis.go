package main

import (
	"fmt"
	"strings"
)

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	elements := []string{"(", ")"}
	var dfs func(int, int, []string)
	dfs = func(left int, right int, path []string) {
		if len(path) == 2*n {
			res = append(res, strings.Join(path, ""))
			return
		}
		for i := 0; i < 2; i++ {
			if (left == right && i == 1) || (left == n && i == 0) || (right == n && i == 1) {
				continue
			}
			var nleft, nright int
			if i == 0 {
				nleft, nright = 1, 0
			} else {
				nleft, nright = 0, 1
			}

			path = append(path, elements[i])
			dfs(left+nleft, right+nright, path)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0, []string{})
	return res
}

func generateParenthesis2(n int) []string {
	res := make([]string, 0)
	var dfs func(int, int, string)
	dfs = func(left int, right int, path string) {
		if len(path) == 2*n {
			res = append(res, path)
			return
		}
		if left < n {
			dfs(left+1, right, path+"(")
		}
		if left > right {
			dfs(left, right+1, path+")")
		}
	}
	dfs(0, 0, "")
	return res
}

func main() {
	fmt.Println(generateParenthesis2(3))
}
