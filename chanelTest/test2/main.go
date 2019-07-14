package main

import (
	// "fmt"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

const (
	url = "https://mp.weixin.qq.com/s/TtiaTA5bDqpAz2VihmI3eg"
)

type T int

// 判断 channel 是否已经关闭了。
func isClosed(ch <-chan T) bool {
	select {
	case <-ch:
		return true
	default:
	}

	return false
}

func sendSafe(ch chan T, value T) (closed bool) {
	defer func() {
		if recover() != nil {
			// 因为已经关闭了，向已经关闭的chan里面发送内容会报错。
			// 所以更改closed 的标志位。
			closed = true
		}
	}()

	ch <- value

	closed = false
	return
}

func SafeClose(ch chan T) (justClosed bool) {
	defer func() {
		if recover() != nil {
			justClosed = false
		}
	}()

	close(ch)
	return true
}

// sync.Once 方式关闭channel
type MyChannel struct {
	C chan T
	sync.Once
	closed bool
	mutex  sync.Mutex
}

func NewMyChannel() *MyChannel {
	return &MyChannel{
		C: make(chan T),
	}
}

func (mc *MyChannel) SafeClose() {
	mc.mutex.Lock()

	mc.Once.Do(func() {
		close(mc.C)
	})

	mc.mutex.Unlock()
}

func (mc *MyChannel) isClosed() (close bool) {
	mc.mutex.Lock()

	defer mc.mutex.Unlock()
	return mc.closed
}

func main() {
	// ch := make(chan T)

	// res := isClosed(ch)
	// if res {
	// 	fmt.Println("res = true")
	// } else {
	// 	fmt.Println("res = false")
	// }

	// close(ch)
	// res = isClosed(ch)
	// if res {
	// 	fmt.Println("res = true")
	// } else {
	// 	fmt.Println("res = false")
	// }

	// ClosePlanA()

	// ClosePlanB()

	// ClosePlan2Example()

	channelBuffer()
}

// 不破坏channel 关闭规则的方案1
// 一个sender 多个receiver
func ClosePlanA() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	const MaxRandomNumber = 10000
	const NumReceivers = 100
	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	dataCh := make(chan int, 10)

	//sender
	go func() {
		for {
			if value := rand.Intn(MaxRandomNumber); value == 0 {
				log.Println("stop value = ", value)
				close(dataCh)
				return
			} else {
				log.Println("value = ", value)
				dataCh <- value
			}
		}
	}()

	// receiver
	for i := 0; i < NumReceivers; i++ {
		go func(i int) {
			defer wgReceivers.Done()

			for value := range dataCh {
				log.Printf("receiver id[%v],  value = %v\n", i, value)
				time.Sleep(100)
			}
		}(i)
	}

	// intSlice := []int{1,2,3,4,5,6}
	wgReceivers.Wait()
}

// 多个sender 一个receiver
func ClosePlanB() {
	const (
		MaxRandomNumber = 10000
		NumSenders      = 100
	)

	// 随机数喂种子。
	rand.Seed(time.Now().UnixNano())

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(1)

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})

	for i := 0; i < NumSenders; i++ {
		go func(i int) {
			for {

				value := rand.Intn(MaxRandomNumber)
				select {
				case <-stopCh:
					log.Printf("sender id = %v, 即将关闭此goroutine， 其他goroutine还在继续写入。", i)
					// close(dataCh)
					return
				case dataCh <- value:
					log.Println("不排除 还有人往里面写 --- ")
				}
			}
		}(i)
	}

	go func() {
		defer wgReceivers.Done()

		for value := range dataCh {
			if value == MaxRandomNumber-1 {
				close(stopCh)
				log.Println("读到了关键的value  关闭 stopCh")
				time.Sleep(time.Second)
				return
			}
		}
	}()

	wgReceivers.Wait()
}

func ClosePlan2Example() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const MaxRandomNumber = 10000
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(1)

	// ...
	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})
	// stopCh is an additional signal channel.
	// Its sender is the receiver of channel dataCh.
	// Its reveivers are the senders of channel dataCh.

	// senders
	for i := 0; i < NumSenders; i++ {
		go func() {
			for {
				value := rand.Intn(MaxRandomNumber)

				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}()
	}

	// the receiver
	go func() {
		defer wgReceivers.Done()

		for value := range dataCh {
			if value == MaxRandomNumber-1 {
				// the receiver of the dataCh channel is
				// also the sender of the stopCh cahnnel.
				// It is safe to close the stop channel here.
				log.Println("读到了关键的value  关闭 stopCh")

				close(stopCh)
				return
			}

			log.Println(value)
		}
	}()

	wgReceivers.Wait()
}

// 测试无缓冲channel 和 缓冲channel
func channelBuffer() {
	intchan := make(chan int)
	intchanBuff := make(chan int, 2)
	wgChan := sync.WaitGroup{}
	wgChan.Add(3)
	count := 0

	// 无缓冲的chan 不可以写入，因为因为没有可以消费intchan 的其他地方。
	go func() {
		for i := 0; i < 10; i++ {
			intchan <- i
			// fmt.Printf("intchan 无缓冲 写入一个值 --- \n")
		}

		wgChan.Done()
		close(intchan)
		count++
		fmt.Println("无缓冲的 chan 写入完毕， close channel ，wg.done()")
	}()

	// 有缓冲的chan 可以写入，但是会报错，因为只有写入，没有读取。会造成死锁。
	go func() {
		for i := 0; i < 10; i++ {
			intchanBuff <- i
			// fmt.Printf("intchanBuff 写入一个值 有缓冲 --- \n ")
		}

		close(intchanBuff)
		wgChan.Done()
		count++

		fmt.Println("有缓冲的 chan 写入完毕， close channel ，wg.done()")
	}()

	go func() {
		for {
			select {
			case value := <-intchan:
				fmt.Println("[读取]intchan value  = ", value)
			case value2 := <-intchanBuff:
				fmt.Println("[读取]intchanBuf value  = ", value2)
			}

			time.Sleep(1000)
			if count == 2 {
				break
			}
		}

		fmt.Println("读取 channel  goroutine 完毕。 0---0")
		wgChan.Done()
	}()

	wgChan.Wait()
}
