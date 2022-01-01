package main

import "reflect"

func f1(data interface{}) {

}

func main() {
	var a int = 1
	_ = reflect.TypeOf(a)

}
