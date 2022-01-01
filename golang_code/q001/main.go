package main

import (
	"fmt"
	"sync"
)

// 交替打印数字和字母
func main() {
	letter, number := make(chan bool), make(chan bool)

	go func() {
		num := 1
		for {
			select {
			case <-number:
				fmt.Print(num)
				num++
				letter <- true
			}
		}
	}()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		le := 'A'
		for {
			select {
			case <-letter:
				fmt.Print(string(le))
				le++
				if le > 'Z' {
					wg.Done()
					return
				}
				number <- true
			}
		}
	}()
	number <- true
	wg.Wait()
}
