package main

import "fmt"

func climbStairs(n int) int {
	//var res int
	//var dfs func(int)
	//dfs = func(gap int) {
	//	if gap <= 0 {
	//		if gap == 0 {
	//			res++
	//		}
	//		return
	//	}
	//	dfs(gap - 1)
	//	dfs(gap - 2)
	//}
	//dfs(n)
	//return res

	if n == 1 {
		return 1
	}
	var first, second, res = 1, 2, 0
	for i := 3; i <= n; i++ {
		res = first + second
		first = second
		second = res

	}
	return res
	//1134903170
}

func main() {
	fmt.Println(climbStairs(44))
}
