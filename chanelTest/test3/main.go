package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

const url = "https://blog.csdn.net/u013862108/article/details/89404966"

// ^^^^^^^^^^^^^^^^^^^^^^^^^ 普通的方式， 去关闭goroutine^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
/*
func printGreeting(done <-chan interface{}) error {
	greeting, err := genGreeting(done)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", greeting)
	return nil
}

func genGreeting(done <-chan interface{}) (string, error) {
	switch locale, err := locale(done); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "hello", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func locale(done <-chan interface{}) (string, error) {
	select {
	case <-done:
		return "", fmt.Errorf("canceled")
	case <-time.After(1 * time.Minute):
	}
	return "EN/US", nil
}

func printFarewell(done <-chan interface{}) error {
	farewell, err := genFarewell(done)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", farewell)
	return nil
}

func genFarewell(done <-chan interface{}) (string, error) {
	switch locale, err := locale(done); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "goodbye", nil
	}
	return "", fmt.Errorf("unsupproted locale")
}

func main() {
	var wg sync.WaitGroup
	done := make(chan interface{})
	defer close(done)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printGreeting(done); err != nil {
			fmt.Println("%v", err)
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printFarewell(done); err != nil {
			fmt.Printf("%v", err)
		}
	}()

	wg.Wait()
}
*/
// vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv

func main() {
	var wg sync.WaitGroup
	//Background() 方法 返回一个空的上下文
	// withCancel 包装它 为了 获取 cancel/取消 方法/功能

	// --------------------------------------------------------------
	// 联系 context 的使用。
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "key", "123")
	res := ctx.Value("key")
	fmt.Println("res = ", res)
	// --------------------------------------------------------------

	runtime.GOMAXPROCS(1)

	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()

		err := printGreeting(ctx)
		if err != nil {
			fmt.Printf("cannot print greeting: %v\n", err)

			//在这一行，如果打印问候出错，将 取消上下文；  或说对上下文 环境 执行 cancel, 这样使用这个上下文或者从这个上下文衍生出来的上下文。
			//中  ct.Done()  将收到信号/收到值； 进而达到解除阻塞 的目的。
			cancel()
			// 这里的cancel 通知了同一个ctx 的 printFarewell（） 中的ctx.Done() locale中的1分钟的阻塞。
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printFarewell(ctx); err != nil {
			fmt.Printf("cannot print farewell: %v\n", err)
		}
	}()
	wg.Wait()

}

func printGreeting(ctx context.Context) error {
	greeting, err := genGreeting(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", greeting)
	return nil
}

func genGreeting(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	switch locale, err := locale(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "hello", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func locale(ctx context.Context) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(1 * time.Minute):
	}
	return "EN/US", nil
}

func printFarewell(ctx context.Context) error {
	farewell, err := genFarewell(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", farewell)
	return nil
}

func genFarewell(ctx context.Context) (string, error) {
	switch locale, err := locale(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "goodbye", nil
	}
	return "", fmt.Errorf("unsupported locale")
}
