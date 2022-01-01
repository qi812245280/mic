package main

import (
	"fmt"
)

func convert2bin(num int) []int {
	if num == 0 {
		return []int{0}
	}
	res := make([]int, 0)
	for num > 0 {
		lsp := num % 2
		res = append([]int{lsp}, res...)
		num = num / 2
	}
	return res
}

func hammingDistance(x int, y int) int {
	//xSlice := convert2bin(x)
	//ySlice := convert2bin(y)
	//if len(xSlice) < len(ySlice) {
	//	xSlice = append(make([]int, len(ySlice)-len(xSlice)), xSlice...)
	//} else if len(ySlice) < len(xSlice) {
	//	ySlice = append(make([]int, len(xSlice)-len(ySlice)), ySlice...)
	//}
	//res := 0
	//for idx, num := range xSlice {
	//	if num != ySlice[idx] {
	//		res += 1
	//	}
	//}
	//return res

	//res := 0
	//n := x ^ y
	//bin := convert2bin(n)
	//for _, num := range bin {
	//	if num == 1 {
	//		res += 1
	//	}
	//}
	//return res

	res := 0
	n := x ^ y
	for n != 0 {
		if n&1 != 0 {
			res += 1
		}
		n = n >> 1
	}
	return res

}

func main() {
	fmt.Println(hammingDistance(1, 4))
}
