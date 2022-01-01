package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
高并发下的锁与map的读写
场景：在一个高并发的web服务器中，要限制IP的频繁访问。现模拟100个IP同时并发访问服务器，每个IP要重复访问1000次。

每个IP三分钟之内只能访问一次。修改以下代码完成该过程，要求能成功输出 success:100

package main

import (
	"fmt"
	"time"
)

type Ban struct {
	visitIPs map[string]time.Time
}

func NewBan() *Ban {
	return &Ban{visitIPs: make(map[string]time.Time)}
}
func (o *Ban) visit(ip string) bool {
	if _, ok := o.visitIPs[ip]; ok {
		return true
	}
	o.visitIPs[ip] = time.Now()
	return false
}
func main() {
	success := 0
	ban := NewBan()
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func() {
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !ban.visit(ip) {
					success++
				}
			}()
		}

	}
	fmt.Println("success:", success)
}
*/

type Ban struct {
	visitIPs map[string]time.Time
	mu       *sync.Mutex
}

func NewBan() *Ban {
	return &Ban{visitIPs: make(map[string]time.Time), mu: &sync.Mutex{}}
}
func (o *Ban) visit(ip string) bool {
	o.mu.Lock()
	defer o.mu.Unlock()
	if v, ok := o.visitIPs[ip]; ok {
		if time.Now().Sub(v) <= 3*time.Minute {
			return true
		}
	}
	o.visitIPs[ip] = time.Now()
	return false
}
func main() {
	var success int32
	ban := NewBan()
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func() {
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !ban.visit(ip) {
					atomic.AddInt32(&success, 1)
				}
			}()
		}

	}
	fmt.Println("success:", success)
}
