package main

import "fmt"

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	var dfs func(map[int]int, []int)
	dfs = func(markMap map[int]int, path []int) {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if markMap[i] == 0 {
				path = append(path, nums[i])
				markMap[i] = 1
				dfs(markMap, path)
				markMap[i] = 0
				path = path[:len(path)-1]
			}
		}
	}
	dfs(map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0}, []int{})
	return res
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute([]int{0, 1}))

}
