package main

import (
	"fmt"

	"github.com/astaxie/goredis"
)

func main() {
	var client goredis.Client
	client.Addr = "127.0.0.1:6379"
	err := client.Set("test", []byte("hello redis"))

	if err != nil {
		fmt.Println("redis set err = ", err)
		return
	}

	value, err := client.Get("test")
	if err != nil {
		fmt.Println("get client value err, err = ", err)
	}
	fmt.Println("value = ", string(value))

	f := make(map[string]interface{})
	f["name"] = "zhangsan"
	f["age"] = 12
	f["sex"] = "女"
	err = client.Hmset("test_hash", f)
	if err != nil {
		panic(err)
	}

}
