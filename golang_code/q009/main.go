package main

import (
	"fmt"
	"sync"
)

/*
	写代码实现两个 goroutine，其中一个产生随机数并写入到 go channel 中，
	另外一个从 channel 中读取数字并打印到标准输出。最终输出五个随机数。
*/

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	go func() {
		defer wg.Done()
		for num := range ch {
			fmt.Println(num)
		}
	}()
	//time.Sleep(100 * time.Second)
	wg.Wait()
}
