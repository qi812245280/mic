package main

import (
	"fmt"
)

func printArray(array []int) {
	for i := range array {
		fmt.Println(array[i])
	}
}

func deferFuncParameter1() {
	var aArray = []int{1, 2, 3}
	defer printArray(aArray)

	aArray[0] = 10
	return
}

func deferFuncParameter2() {
	var aInt = 1
	defer func(i *int) {
		fmt.Println(*i)
	}(&aInt)
	aInt = 2
	return
}

func deferFuncReturn() (result int) {
	i := 1
	defer func() {
		result++
	}()
	return i
}

func foo() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func main() {
	fmt.Println(foo())
	//json.Marshal()
	//http.ListenAndServe()
}
