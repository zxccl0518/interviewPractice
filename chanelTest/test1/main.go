package main

import (
	"fmt"
	"sync"
	"time"
)

type ThreadSafeSet struct {
	sync.RWMutex
	s []int
}

func (t *ThreadSafeSet) Iter() <-chan interface{} {
	m := make(chan interface{})
	go func() {
		t.RLock()

		for v := range t.s {
			m <- v
			fmt.Println("[写] Iter v = ", v)
			// 这个位置读取不到任何内容,因为阻塞了, 创建的m channe 是一个无缓冲的通道, 当没有接收者的时候,会一直处于阻塞状态.
			// m := make(chan interface{}) 不同于 m:=make(chan interface{}, 1) ,后者是一个有缓冲的通道, 缓冲为1, 没有接受者,也不会阻塞,
			// 当已经写进一个内容并且 在没有接收者读取的时候,再写进第二个内容会发生阻塞. 这就是区别.
		}

		close(m)
		t.RUnlock()
	}()

	return m
}

func read() {
	th := ThreadSafeSet{}
	th.s = make([]int, 100)
	ch := th.Iter()
	close := false

	for {

		select {
		case v, ok := <-ch:
			if ok == false {
				close = true
			} else {
				fmt.Println("[读] read =", v)
				time.Sleep(time.Second * 3)
			}
		}

		if close {
			break
		}
	}

	fmt.Println("read done===========")
}

func unread() {
	th := ThreadSafeSet{}
	th.s = make([]int, 100)
	ch := th.Iter()
	_ = ch
	time.Sleep(time.Second * 5)
	fmt.Println("unread done --- ")
}

func main() {
	read()

	// unread()
}
