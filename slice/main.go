package main

import (
	"fmt"
	// "encoding/json"
)


func main(){
	
	
	
	// 这个主要是查看 json标准库对 nil slice 和 空slice 的区别。 但是好像没有查看出来。
	/*
	type Animal struct { 
		Name  string 
		Order string 
		} 
		var jsonBlob = [ ] byte ( ` [ 
			{ "Name" : "Platypus" , "Order" : "Monotremata" } , 
			{ "Name" : "Quoll" ,     "Order" : "Dasyuromorphia" } 
			] ` ) 
			
			//  空slice 的情況下，進行反序列化。
			// var sliceEmp = []Animal{}
			
			// nil slce 情況下，進行反序列化
	var sliceNil []Animal
			
	// err:=json.Unmarshal(jsonBlob,&sliceEmp)
	err:=json.Unmarshal(jsonBlob,&sliceNil)
	if err != nil{
		fmt.Println("反序列化失败。err =  ",err)
	}else{
		fmt.Println("反序列化 成功。。。。")
	}
	*/

	// 考点是 slice nil 和 empty 的问题
	// var sliceNil []int
	// sliceNil[0] = 1
	// fmt.Println("sliceNil = ",sliceNil)
	//panic: runtime error: index out of range 报错了。 越界了。
	// sliceNil = append(sliceNil,1)
	// fmt.Println("sliceEmpty = ",sliceNil)


	// var sliceEmpty  = []int{}
	// sliceEmpty[0] = 1
	// fmt.Println("sliceEmpty = ",sliceEmpty)
	//panic: runtime error: index out of range 报错了。 也会提示 越界了。因为初始化的时候 就是默认0个数据。如果想增加数据 用append
	// sliceEmpty = append(sliceEmpty,1)
	// fmt.Println("sliceEmpty = ",sliceEmpty)
	// 这个时候是没问提的。



	/*
	url := "https://blog.csdn.net/lolimostlovely/article/details/80717701"
	// slice 的拷贝问题
	// 目前知道的有2种方式拷贝。 一个是浅拷贝， 一个是深拷贝。
	// 浅拷贝就是 一个切片直接拷贝另一个 切片，但是这个浅拷贝的方式有个问题 就是切片是引用，地址拷贝， 当其中的一个切片改动了， 会影响其他的浅拷贝的切片。
	// 深拷贝 ，不是直接引用切片的地址， 而是需要在拷贝之前 提前申请一个新的slice 然后将一个slice的内容拷贝到新的slice中。
	// eg：
	// slice := make([]int, 5,5)
	// slice1:= slice
	// slice2 := slice[:]
	// slice3 := slice[0:4]
	// slice4 := slice[1:5]
	// slice[1] = 1
	// fmt.Println("slice",slice)
	// fmt.Println("slice1",slice1)
	// fmt.Println("slice2",slice2)
	// fmt.Println("slice3",slice3)
	// fmt.Println("slice4",slice4)
	// 可以看出来， 浅拷贝的方式 拷贝slice 会有一个改动，影响其他的slice 的问题。


	// 接下来看 深拷贝的使用。
	slice := make([]int, 5,5)
	slice1 := slice
	slice = append(slice,1)	
	slice[0] = 1
	// 因为append 这个操作是 操作的底层的匿名数组， 因为slice初始化的时候是5的最大容量，
	// 现在append 的时候，容量不够了， 会重新再底层开一个新的匿名数组，将原来的只拷贝过来，所以底层的数组的地址已经发生了改变，
	// 所以 在 appendK 扩容之后，在对原来的slice[0]=1 的这种改变不会影响新的slice1了。

	// 但是如果 slice[0]=1是在扩容之前的化， slice 和 slice1 的值都会改变。
	fmt.Println("slice = ",slice)
	fmt.Println("slice1 = ",slice1)

	slice2 := make([]int, 4,4)
	slice3 := make([]int, 5,5)
	fmt.Println("slice = ",slice)
	fmt.Println("slice2 = ", copy(slice2,slice)) // 拷贝了4个值，
	fmt.Println("slice2 = ", copy(slice3,slice)) // 拷贝了5个值。
	
	slice2[1] = 2
	slice3[1] = 3
	fmt.Println("--------------------------------------")
	fmt.Println("slice = ",slice)
	fmt.Println("slice2 = ",slice2)
	fmt.Println("slice3 = ",slice3)
	// copy 深拷贝之后，slice 的改动不会影响其他的slice
	*/
}


