package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	minValue := math.MaxInt64
	var maxPro int
	for _, price := range prices {
		if price < minValue {
			minValue = price
		}
		if price-minValue > maxPro {
			maxPro = price - minValue
		}
	}
	return maxPro
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
