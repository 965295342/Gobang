package main

import (
	"GoGame/def"
	mynetwork "GoGame/network"
	"GoGame/server"
	"fmt"
	"log"
	"net"

	"google.golang.org/protobuf/proto"
)

func main() {
	myaddress, err := mynetwork.GetLocalIPV4()
	if err != nil {
		log.Println(err)
	}
	log.Println(myaddress)

	def.MYIP = "127.0.0.1"
	SocketService, err := server.NewSocketService(":8848")
	if err != nil {
		log.Println(err)
	}
	testsSend()
	for {
		//等待客户的连接?
		conn, err := SocketService.Listener.Accept()
		//如果有错�?直接跳过
		if err != nil {
			continue
		}

		//通过goroutine协程处理连接
		go readHandle(conn)
		// writeHandle(conn)

	}

}

func readHandle(conn net.Conn) {
	defer conn.Close()
	for {
		// 接收数据
		buf := make([]byte, 1024)
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		if n == 0 {
			continue
		}
		rAddr := conn.RemoteAddr()

		test := def.NormalMessageC2S{}
		msg := proto.Unmarshal(buf, &test)

		fmt.Println("Receive from client", rAddr.String(), test, n)

		if msg != nil {
			fmt.Println(msg)
		}
		messageHandler(&test)
		break

	}
}

// func writeHandle(conn net.Conn) {
// 	defer conn.Close()
// 	host := ":8849"
// 	var data []byte
// 	for {
// 		if len(server.SendStack) > 0 {
// 			data, server.SendStack = server.SendStack[len(server.SendStack)-1], server.SendStack[:len(server.SendStack)-1]

// 			myAddress := def.MYIP + host
// 			conn, err := net.Dial("tcp", myAddress)
// 			if err != nil {
// 				fmt.Errorf("连接失败:", err)
// 			}

// 			n, err := conn.Write(data)
// 			if err != nil {
// 				fmt.Println("发送数据失败")
// 				return
// 			}
// 			fmt.Printf("一共发送了%d个字节的数据\n", n)
// 		}
// 	}
// }
func messageHandler(pack *def.NormalMessageC2S) {
	hanlder := server.CallBackMap[pack.ID]
	if hanlder == nil {
		return
	}
	hanlder(pack)
}

func testsSend() {
	newMessage := def.NormalMessageS2C{}
	newMessage.STRING = "Hello,world"
	data, err := proto.Marshal(&newMessage)
	if err != nil {
		log.Print(err)
	}

	server.Send(def.HEART_BEAT, data)
}
