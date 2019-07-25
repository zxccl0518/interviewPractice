package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func Cli_MessageSend(conn net.Conn) {
	var sendInfo string
	for {
		reader := bufio.NewReader(os.Stdin)
		sendByte, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("输入有误， 不得发送。")
			continue
		}

		// 注意 出现意外要关闭conn链接
		if strings.ToUpper(string(sendByte)) == "EXIT" {
			conn.Close()
			break
		}

		// fmt.Println("查看client 本地的ip = ", conn.LocalAddr().String())
		sendInfo = conn.LocalAddr().String() + "#" + string(sendByte)
		_, err = conn.Write([]byte(sendInfo))
		if err != nil {
			conn.Close()
			fmt.Println("client 发送消息失败")
		}
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic("链接服务器 失败。====")
	}
	defer conn.Close()

	go Cli_MessageSend(conn)

	var recBuff = make([]byte, 1000)
	for {
		readNums, err := conn.Read(recBuff)
		if err != nil {
			fmt.Println("客户端 读取server 返回的message 有误。")
			continue
		}
		fmt.Printf("server:%v\n", string(recBuff[:readNums]))
	}
}
