package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func main()  {
//     defer func() {
//         if err:=recover();err!=nil{
//             fmt.Println("++++")
//             f:=err.(func()string)
//             fmt.Println("err = ",err,", --- f() = ",f(),", --- reflect = ",reflect.TypeOf(err).Kind().String())
//         }else {
//             fmt.Println("fatal")
//         }
//     }()

//     defer func() {
//         panic(func() string {
//             return  "defer panic"
//         })
//     }()
//     panic("panic")
// }


// import (
// 	"fmt"
// 	"strings"
// )

// func main(){
// 	fmt.Println(Utf8Index("北京天安门最美丽", "天安门"))
// 	fmt.Println(strings.Index("北京天安门最美丽", "男"))
// 	fmt.Println(strings.Index("", "男"))
// 	fmt.Println(Utf8Index("12ws北京天安门最美丽", "天安门"))
// }

// func Utf8Index(str, substr string) int {
// 	asciiPos := strings.Index(str, substr)
// 	if asciiPos == -1 || asciiPos == 0 {
// 		return asciiPos
// 	}
// 	pos := 0
// 	totalSize := 0
// 	reader := strings.NewReader(str)
// 	for _, size, err := reader.ReadRune(); err == nil; _, size, err = reader.ReadRune() {
// 		totalSize += size
// 		pos++
// 		// 匹配到
// 		if totalSize == asciiPos {
// 			return pos
// 		}
// 	}
// 	return pos
// }




import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}

		fmt.Println("发送端 关闭channel --- ")
		close(ch)
	}()
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	// close(ch)
	fmt.Println("bottom ok ~~~")
	time.Sleep(time.Second * 100)
}



