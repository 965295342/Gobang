package main

import (
	"GoGame/def"
	mynetwork "GoGame/network"
	"GoGame/server"
	"fmt"
	"log"
	"net"

	gophp "github.com/techoner/gophp"
)

func main() {
	myaddress, err := mynetwork.GetLocalIPV4()
	if err != nil {
		log.Println(err)
	}
	log.Println(myaddress)

	def.MYIP = myaddress
	SocketService, err := server.NewSocketService()
	if err != nil {
		log.Println(err)
	}

	for {
		//等待客户的连�?
		conn, err := SocketService.Listener.Accept()
		//如果有错�?直接跳过
		if err != nil {
			continue
		}
		//通过goroutine来�?�理用户的�?�求
		go clientHandle(conn)
	}

}

func clientHandle(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		// 接收数据
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		rAddr := conn.RemoteAddr()
		msg := string(buf[0:n])
		fmt.Println("Receive from client", rAddr.String(), msg)

		out, _ := gophp.Unserialize([]byte(buf))
		fmt.Println("Receive from client", rAddr.String(), out)
		// var m Msg
		// err = json.Unmarshal([]byte(msg), &m)
		// if (err != nil || m == Msg{}) {
		// 	conn.Write([]byte("1"))
		// 	return
		// }

		// if m.T == "c"{
		// 	handleConsumer(conn, m)
		// }else {
		// 	handleProduct(conn, m)
		// }
	}
}
