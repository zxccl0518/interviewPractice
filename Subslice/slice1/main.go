package main

import (
	"bytes"
	"container/heap"
	"context"
	"fmt"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

func sliceTest1() {
	s := []int{1, 2, 3}
	// s := make([]int, 10)
	ss := s[1:]

	// append 一旦对slice 进行扩容了, 那么返回的的slice 与原来的slice 地址引用就不同了.
	// 也就是说是新的 slice ,彼此改动不会影响到彼此.
	ss = append(ss, 4)

	// 说明通过这个range 的方式改变slice 是错的,
	// v 就是一个临时变量, 这个变量里面 不停的复制 slice 的值.
	for _, v := range ss {
		v += 10
	}

	fmt.Println(" s = ", s)
	fmt.Println(" ss = ", ss)

	// 要想改变这个slice的值, 要通过下表索引的方式 去改变值.
	for i := range ss {
		ss[i] += 10
	}

	fmt.Println("--------------------------")
	fmt.Println(" s = ", s)
	fmt.Println(" ss = ", ss)
}

// ---------------------------------------
type S struct {
}
type IF interface {
	F()
}

func (s S) F() {

}

func InitType() S {
	var s S
	return s
}
func InitPointer() *S {
	var s *S
	return s
}

func InitEfaceType() interface{} {
	var s S
	return s
}

func InitEfacePointer() interface{} {
	var s *S
	return s
}

func InitIfaceType() IF {
	var s S
	return s
}

func InitIfacePointer() IF {
	var s *S
	return s
}

func Interface_structTest1() {
	// 无法通过编译, 因为无法将 nil 转换为 S
	// cannot convert nil to type Sgo
	// fmt.Println(InitType() == nil)

	// 返回的是结构体的指针,因为结构体没有实例化,只有声明,所以指针 是一个nil的空值.所以返回true
	fmt.Println(InitPointer() == nil)
	// 因为这个函数返回值是 interface{}, 所以return之前 会将变量变成一个空的interface{} 但是空interface 是数据为空,类型不为空.
	// 所以 != nil  返回false
	fmt.Println(InitEfaceType() == nil)
	// 同理 返回的一个空接口 != nil  结果为false
	fmt.Println(InitEfacePointer() == nil)

	// 返回值  是一个已经被实现了的接口, 不为空.
	fmt.Println(InitIfaceType() == nil)
	fmt.Println(InitIfacePointer() == nil)
}

//--------------------

type V struct {
	m string
}

func f() *V {
	return &V{"foo"} //A
}

func returnTest1() {
	p := *f()  //B
	print(p.m) //print "foo"
}

//-----------------------------------------------------------
type S1 struct {
}

func z(x interface{}) {
}

func g(x *interface{}) {
}

func InterfaceTest2() {
	s := S1{}
	p := &s
	var w interface{}
	z(s) //A
	// g(s) //B  cannot use s (type S1) as type *interface {} in argument to g:*interface {} is pointer to interface, not interface
	z(p) //C
	// g(p) //D  cannot use s (type *S1) as type *interface {} in argument to g:*interface {} is pointer to interface, not interface

	// 下面2种情况 是对的.
	w = s
	g(&w) // E
	w = p
	g(&w) // F
}

//-----------------------------------------------------------
// Add code in line A to assure that the lowercase letters and capital letters are printed consecutively.
// 看明白了 这个 Gosched()这个函数的意思了. 就是吧当前的cpu的执行权让出来,但是不会将goroutine挂起.
// 让出执行权的goroutine 会自动被恢复的.
func GoSchedTest() {
	const GOMAXPROCS = 1
	const N = 26
	runtime.GOMAXPROCS(GOMAXPROCS)

	var wg sync.WaitGroup
	wg.Add(2 * N)
	for i := 0; i < N; i++ {
		go func(i int) {
			defer wg.Done()
			// A
			runtime.Gosched()
			fmt.Printf("%c", 'a'+i)
		}(i)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("%c", 'A'+i)
		}(i)
	}
	wg.Wait()
}

//-----------------------------------------------------------

// 这个题目的考点是 for 循环10次， 一瞬间就已经完成了，因为i没有作为参数传进来，所以剩下10个线程 共同访问一个变量数据 i， 而不是10个不同值，不同地址的i
// 但是这个i变量资源 被互斥锁 mutex 保护起来， 访问这个资源 就要竞争，抢锁的资源
// 此时 这个for循环中的变量i 是一个变量，循环期间被赋予了不同的值，但是地址只有一个。
// 所以循环结束的时候， i已经是10了，10个协程抢到锁的资源的时候，拿到的i 的值是10 。
// map 的key 是同一个值10 所以map 的长度也就是1 改进方式是 将i作为参数 传进去，这样就有了值拷贝，各个协程访问到的资源就是当时传入参数的值。
func mutexTest() {
	N := 10
	m := make(map[int]int)

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("获取 mutex 互斥锁之前 --- ")
			mu.Lock()
			fmt.Println("获取到了 mutex 互斥锁 ---  i = ", i)

			m[i] = i
			mu.Unlock()
		}()
	}
	wg.Wait()
	for i, v := range m {
		fmt.Printf("i = %v, v = %v\n", i, v)
	}
	fmt.Println(len(m))
}

//-----------------------------------------------------------
// Map Immutability
// 没太明白 这个map 的原理.
type mapS struct {
	name string
}

func MapTest() {
	//cannot assign to struct field m["x"].name in map
	m := map[string]mapS{"x": mapS{"one"}}
	// m["x"].name = "two"

	// 正确写法,
	// m := map[string]*mapS{"x": &mapS{"one"}}
	// m["x"].name = "two"
	fmt.Println("m = ", m)
}

//-----------------------------------------------------------
//Add code to line A to sort s in ascending order
// 考点是 对于slice 的排序, 可以用 sort.Slice()这个方法进行排序.
// sort.SliceStable() 这个方法是相对于sort.Slice() 稳定的方法..
type sliceS struct {
	v int
}

func SliceSortingTest() {
	s := []sliceS{{1}, {3}, {5}, {2}}
	//A
	sort.SliceStable(s, func(i, j int) bool { return s[i].v < s[j].v })
	fmt.Printf("%#v ", s)
}

//-----------------------------------------------------------
// utf8 length
//Fix the mistake below to assure the length of utf8 string can be printed correctly.
func Utf8Test() {
	// len() 返回的是字节数.  一个汉字3个字节.
	fmt.Println(len("你好"))

	// 若想返回的是直接的 汉字的个数 可以 用 utf8.RuneCountString()方法.
	fmt.Println(utf8.RuneCountInString("你好"))
	fmt.Println(utf8.RuneCountInString("hello"))

	fmt.Println(strings.IndexRune("helloworld", 'l'))
	fmt.Println(strings.IndexRune("今晚打老虎", '打'))

	fmt.Println("bytes.Count() = ", bytes.Count([]byte("abc123今晚打老虎"), []byte("打")))
	fmt.Println("utf8.RuneCount() = ", utf8.RuneCount([]byte("abc123今晚打老虎")))
	fmt.Println("strings.Count() = ", strings.Count("abc123今晚打老虎", "打"))
}

//-----------------------------------------------------------
// What will be printed when the code below is executed?
func HeapTest() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//-----------------------------------------------------------
//What will be printed when the code below is executed?
// 结论是 不会等到3秒的时候 在让ctx接收到 done（）的信号， 等待1秒的时候就放回了。
func contextWithTimeOut() {
	timeout := 3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("waited  for 1 sec")
	case <-time.After(1 * time.Second):
		fmt.Println("waited  for 2 sec")
	case <-time.After(1 * time.Second):
		fmt.Println("waited  for 3 sec")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

//-----------------------------------------------------------
// append之后 扩容了， 原来底层维护的数组不再使用。所以append之后的slice 地址就变了。
func sliceAddress() {
	slice1 := make([]int, 10)
	fmt.Printf("地址 ：%p\n", slice1)
	for v := range slice1 {
		fmt.Println("v = ", v)
	}

	slice1 = append(slice1, 1)
	fmt.Printf("地址 ：%p\n", slice1)

	var i = -26
	fmt.Printf("i = %08b\n", i)
	fmt.Println("i = ", i)
	i = i >> 4
	fmt.Printf("i = %08b\n", i)
	fmt.Println("i = ", i)

}

//-----------------------------------------------------------
//
// https://blog.csdn.net/wowzai/article/details/8941865
func UnicodeIndex(str, substr string) int {
	// 子串在字符串的字节位置
	result := strings.Index(str, substr)
	if result >= 0 {
		// 获得子串之前的字符串并转换成[]byte
		prefix := []byte(str)[0:result]
		// 将子串之前的字符串转换成[]rune
		rs := []rune(string(prefix))
		// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
		result = len(rs)
	}

	return result
}

//-----------------------------------------------------------

func main() {
	// sliceTest1()

	// Interface_structTest1()

	// returnTest1()

	// InterfaceTest2()

	// mutexTest()

	// GoSchedTest()

	// MapTest()

	// SliceSortingTest()

	Utf8Test()

	// 应该是测试 go heap堆得一个 包里面的方法。
	// HeapTest()

	// contextWithTimeOut()

	// sliceAddress()

	// fmt.Println(UnicodeIndex("abc123今晚打老虎", "打"))
}
