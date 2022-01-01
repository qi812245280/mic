package main

import "fmt"

func countBits(n int) []int {
	res := []int{0}
	for i := 1; i <= n; i++ {
		num := i
		cnt := 0
		for num > 0 {
			if num&1 == 1 {
				cnt += 1
			}
			num = num >> 1
		}
		res = append(res, cnt)
	}
	return res
}

func main() {
	fmt.Println(countBits(5))
}
