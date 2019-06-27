package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

type chanObject struct {
	channel chan int
	count   int
}

// 生产者, 每个一秒,生产一个
func (ci chanObject) Add() {
	fmt.Println("开始 Add ------------Add---------")
	wg.Add(1)

	go func() {
		for i := 0; i < 20; i++ {
			ci.channel <- i
			fmt.Println("生产  i = ", i)
			if i == 19 {
				wg.Done()
				close(ci.channel)
			}
		}
	}()
}

// 消费者.
// func (ci chanObject) Sub() {
// 	wg.Add(1)

// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			var res = <-ci.channel
// 			fmt.Printf("消费者 res = %v ", res)
// 			time.Sleep(time.Second * 1)
// 		}

// 		wg.Done()
// 	}()
// }

func (ci chanObject) Sub() {
	fmt.Println("开始 sub ------------sub---------")
	wg.Add(1)

	go func() {
		for {
			select {
			case v, ok := <-ci.channel:
				if ok == false {
					fmt.Println("管道已经关闭了.----")
					wg.Done()

				}

				fmt.Println("<-ci.channel -=----------11---------v = ", v)
				ci.count = v
				fmt.Println("<-ci.channel -=---------22---------- count = ", ci.count)

				break
			}

			time.Sleep(time.Second)
		}
	}()
}

func main() {
	//-----=--=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
	// 有缓冲的通道, 当通道已经满了,再写入会出现阻塞情况.
	// 如果 当前通道已经空了,且检测到会有新的内容写进通道,那么当前的读操作就会产生阻塞.
	wg = sync.WaitGroup{}

	cb := chanObject{
		channel: make(chan int, 10),
		count:   0,
	}

	// 调用消费者.
	cb.Sub()

	// 调用生产者
	cb.Add()

	wg.Wait()
	//-----=--=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

	//--------------------------------------------------------------
	// 通过这个例子,说明
	// 1)无缓冲的channel 当从一个空的channel读取的时候回产生阻塞,直到 有内容写进通道.
	// 2)无缓冲的channel 向一个满的channel 写数据的时候,会产生阻塞,直到有其他位置 读取通道的内容为止.
	// chanint := make(chan int)
	// go func() {
	// 	time.Sleep(5 * time.Second)
	// 	chanint <- 1
	// }()

	// <-chanint
	// fmt.Println("_________________看看是不是阻塞____________________")
	//--------------------------------------------------------------

	// 注意,想已经关闭的通道中发送内容的时候 会出现panic
	// 从一个已经关闭的通道中 读取数据, 返回2个值, 第一是零值, 第二个是false
}
