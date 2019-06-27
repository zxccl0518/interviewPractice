package main

import (
	"fmt"
	"time"
)

func main() {
	// 向一个 值为nil的chan中 读\写 都会造成永久的阻塞,  利用这个特性可以制作select的动态打开和关闭case语句.
	var inCh = make(chan int)
	var outCh = make(chan int)

	go func() {
		var in <-chan int = inCh
		var out chan<- int
		var val int

		for {
			select {
			case out <- val:
				// time.Sleep(time.Nanosecond)				这里做延迟处理的话, 就相当于在执行out = nil之前有一段空挡时间,就会在out阻塞之前 有一定几率 被协程2 抢占执行.打印出outch中的i2
				fmt.Println("--------1--------")
				out = nil
				fmt.Println("--------2--------")
				in = inCh
				fmt.Println("--------3--------")
			case val = <-in:
				fmt.Println("+++++++++1+++++++")
				in = nil
				fmt.Println("+++++++++2+++++++")
				out = outCh
				time.Sleep(time.Nanosecond) // 如果这里 做了延迟处理, 就会有一定几率在out = outch之后,阻塞的outch 变成非阻塞, 被线程2 抢占了outch资源,将outch打印出来,但是值能打印1 打印结束之之后,循环结束 就不会有机会打印 outch 2的结果了.
				fmt.Println("+++++++++3+++++++")
			}
		}
	}()

	go func() {
		for v := range outCh {
			fmt.Println("Result = ", v)
		}
		// for {
		// 	fmt.Println("第二个协程  监听 outch的变化, <-outch = ", <-outCh)
		// }
	}()

	time.Sleep(0)
	inCh <- 1
	inCh <- 2
	time.Sleep(3)
}
