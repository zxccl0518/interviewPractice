package main

import (
	"fmt"
	"net"
	"strings"
)

var messageChan = make(chan string, 1000)
var stopCh = make(chan struct{})
var onlineUsers = make(map[string]net.Conn)

func ProcessInfo(conn net.Conn) {
	defer func(conn net.Conn) {
		// 服务器接收客户端数据异常的话， 关闭这个对应的客户端的tcp链接。
		conn.Close()

		// 维护在线用户列表。
		delete(onlineUsers, conn.RemoteAddr().String())

	}(conn)

	buf := make([]byte, 1024)

	for {
		numOfBytes, err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务器读取 客户端发送来的消息失败")
			break
		}

		if numOfBytes != 0 {
			content := string(buf[:numOfBytes])

			//将消息放入到消息队列
			messageChan <- content
			// fmt.Printf("addr = %s, message = %s\n", conn.RemoteAddr().String(), string(content))

			// 服务器收到了客户端的消息， 回复了客户端一句话。 这句话就是在客户端发过来的消息前面加个前缀， server say
			// conn.Write([]byte("server say : " + content))
		}
	}
}

// 消费者
func consumerProcess() {
	for {
		select {
		case value := <-messageChan:
			parseMessage(value)
			// case <-stopCh:
			// 	break
		}
	}
}

func parseMessage(value string) {
	contentArr := strings.Split(value, "#")
	fmt.Println("contentArr = ", contentArr, " len contentArr = ", len(contentArr))

	if len(contentArr) > 1 {
		addr := contentArr[0]
		strings.Trim(addr, " ")

		if conn := onlineUsers[addr]; conn != nil {
			sentence := strings.Join(contentArr[1:], "#")

			_, err := conn.Write([]byte(sentence))
			if err != nil {
				fmt.Println("服务器 发送消息失败。")
			}
		} else {
			fmt.Println("没有检测到有效的 用户ip地址。 无法回复消息。")
		}
	} else {
		// 其他功能。
	}
}

func ListOnlineUsers() {
	for k := range onlineUsers {
		fmt.Printf("user ip = [%v] \n", k)
	}
}

func main() {
	listen_socket, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic("net listen failed")
	}
	defer listen_socket.Close()

	go consumerProcess()

	for {
		conn, err := listen_socket.Accept()
		if err != nil {
			panic("accept failed")
		}

		onlineUsers[conn.RemoteAddr().String()] = conn

		// 每个用户登录，服务器端都会list 一遍所有在线的用户.
		ListOnlineUsers()

		go ProcessInfo(conn)
	}
}
