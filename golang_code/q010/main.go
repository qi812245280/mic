package main

import (
	"fmt"
	"sync"
	"time"
)

/*
GO里面MAP如何实现key不存在 get操作等待 直到key存在或者超时，保证并发安全，且需要实现以下接口：

type sp interface {
    Out(key string, val interface{})  //存入key /val，如果该key读取的goroutine挂起，则唤醒。
										此方法不会阻塞，时刻都可以立即执行并返回
    Rd(key string, timeout time.Duration) interface{}  //读取一个key，如果key不存在阻塞，等待key存在或者超时
}
*/
type entry struct {
	Value   interface{}
	Ch      chan struct{}
	IsExtra bool
}

type Map struct {
	data map[string]*entry
	lock *sync.RWMutex
}

func (m *Map) Out(key string, val interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if v, ok := m.data[key]; !ok {
		newE := &entry{
			Value:   val,
			Ch:      make(chan struct{}),
			IsExtra: true,
		}
		m.data[key] = newE
		newE.Ch <- struct{}{}
	} else {
		v.IsExtra = true
		v.Value = val
		v.Ch <- struct{}{}
	}
}

func (m *Map) Rd(key string, timeout time.Duration) interface{} {
	m.lock.Lock()
	if v, ok := m.data[key]; ok {
		if v.IsExtra {
			m.lock.Unlock()
			return v.Value
		} else {
			m.lock.Unlock()
			fmt.Println("阻塞中...")
			select {
			case <-v.Ch:
				return v.Value
			case <-time.After(timeout):
				fmt.Println("阻塞超时")
				return nil
			}
		}
	}
	newE := &entry{
		Value:   nil,
		Ch:      make(chan struct{}),
		IsExtra: false,
	}
	m.data[key] = newE
	m.lock.Unlock()
	fmt.Println("阻塞中...")
	select {
	case <-newE.Ch:
		return newE.Value
	case <-time.After(timeout):
		fmt.Println("阻塞超时")
		return nil
	}
}

func main() {
	m := Map{data: make(map[string]*entry), lock: &sync.RWMutex{}}
	go func() {
		fmt.Println(m.Rd("qi", 10*time.Second))
	}()
	time.Sleep(1 * time.Second)
	m.Out("qi", "哎呀呀")
	time.Sleep(10 * time.Second)
}
