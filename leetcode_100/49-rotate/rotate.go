package main

import "fmt"

func rotate(matrix [][]int)  {
	for i := 0 ; i <len(matrix); i++ {
		for j := i+1; j <len(matrix); j ++ {
			matrix[i][j],matrix[j][i]  = matrix[j][i], matrix[i][j]
		}
	}
	for _, nums := range matrix{
		for i:= 0;i < len(nums)/2; i++ {
			nums[i], nums[len(nums) - 1-i] = nums[len(nums) - 1-i], nums[i]
		}
	}
}


func main()  {
	//ori := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	ori := [][]int{{5,1,9,11},{2,4,8,10},{13,3,6,7},{15,14,12,16}}
	rotate(ori)
	fmt.Println(ori)
}
