package server

import (
	"GoGame/def"
	"fmt"
	"net"
	"sync"
	"time"
)

var CallBackMap map[int32]func(*def.NormalMessageC2S)
var SendStack [][]byte

func init() {
	CallBackMap = make(map[int32]func(*def.NormalMessageC2S))
	if CallBackMap == nil {
		fmt.Errorf("CallBack Init Error")
	}
	SendStack = make([][]byte, 0)
	def.MYIP = "127.0.0.1"
}

func Register(messageType int32, hanlder func(*def.NormalMessageC2S)) {
	if CallBackMap[messageType] != nil {
		fmt.Println("Message has registered", messageType)
	} else {
		CallBackMap[messageType] = hanlder
		fmt.Println("Message register success , Message:", messageType)
	}
}

func NewSocketService(host string) (*def.SocketService, error) {
	l, err := net.Listen("tcp", host)
	if err != nil {
		return nil, err
	}
	s := &def.SocketService{
		Sessions:   &sync.Map{},
		StopCh:     make(chan error),
		HbInterval: 0 * time.Second,
		HbTimeout:  0 * time.Second,
		Laddr:      host,
		Status:     def.STInited,
		Listener:   l,
	}

	return s, nil
}

func Send(MessageType int32, data []byte) {
	host := ":8849"
	myAddress := def.MYIP + host
	conn, err := net.Dial("tcp", myAddress)
	if err != nil {
		fmt.Errorf("连接失败:", err)
	}

	n, err := conn.Write(data)
	if err != nil {
		fmt.Println("发送数据失败")
		return
	}
	fmt.Printf("一共发送了%d个字节的数据\n", n)
	SendStack = append(SendStack, data)
}

// buf := bytes.Buffer{}
// encoder := gob.NewEncoder(&buf)

// err := encoder.Encode(p)
// if err != nil {
// 	fmt.Println("编码失败,错误原因: ", err)
// 	return
// }
