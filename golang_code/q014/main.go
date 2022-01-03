package main

/*
写出以下代码出现的问题
package main
import (
    "fmt"
)
func main() {
    var x string = nil
    if x == nil {
        x = "default"
    }
    fmt.Println(x)
}
golang 中字符串是不能赋值 nil 的，也不能跟 nil 比较。
*/

/*
const (
	a = iota
	b = iota
)
const (
	name = "menglu"
	c    = iota
	d    = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
*/

/*
type query func(string) string

func exec(name string, vs ...query) string {
	ch := make(chan string)
	fn := func(i int) {
		ch <- vs[i](name)
	}
	for i, _ := range vs {
		go fn(i)
	}
	return <-ch
}

func main() {
	ret := exec("111", func(n string) string {
		return n + "func1"
	}, func(n string) string {
		return n + "func2"
	}, func(n string) string {
		return n + "func3"
	}, func(n string) string {
		return n + "func4"
	})
	fmt.Println(ret)
}

上面的代码有严重的内存泄漏问题，出错的位置是 go fn(i)，
实际上代码执行后会启动 4 个协程，但是因为 ch 是非缓冲的，
只可能有一个协程写入成功。而其他三个协程会一直在后台等待写入。
*/

/*
func main() {
	str1 := []string{"a", "b", "c"}
	str2 := str1[1:]
	str2[1] = "new"
	fmt.Println(str1)
	str2 = append(str2, "z", "x", "y")
	fmt.Println(str1)
}
*/

/*
type Student struct {
	Name string
}

func main() {
	fmt.Println(&Student{Name: "menglu"} == &Student{Name: "menglu"})
	fmt.Println(Student{Name: "menglu"} == Student{Name: "menglu"})
}
*/

/*
func main() {
	fmt.Println([...]string{"1"} == [...]string{"1"})
	fmt.Println([]string{"1"} == []string{"1"})
}
*/

/*
type Student struct {
	Age int
}

func main() {
	kv := map[string]Student{"menglu": {Age: 21}}
	kv["menglu"].Age = 22
	s := []Student{{Age: 21}}
	s[0].Age = 22
	fmt.Println(kv, s)
}
*/
