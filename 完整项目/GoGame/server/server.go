package server

import (
	"GoGame/def"
	"fmt"
	"net"
	"sync"
	"time"
)

var CallBackMap map[int32]func(*def.Message)

type Player struct {
	id int32
}

func init() {
	CallBackMap = make(map[int32]func(*def.Message))
	if CallBackMap == nil {
		fmt.Errorf("CallBack Init Error")
	}
}

func Register(messageType int32, hanlder func(*def.Message)) {
	if CallBackMap[messageType] != nil {
		fmt.Println("Message has registered", messageType)
	} else {
		CallBackMap[messageType] = hanlder
		fmt.Println("Message register success , Message:", messageType)
	}
}

func NewSocketService() (*def.SocketService, error) {
	var host string = ":8848"
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
	var host string = ":8848"
	conn, err := net.Dial("tcp", def.MYIP+host)
	if err != nil {
		fmt.Errorf("连接失败:", err)
	}

	n, err := conn.Write(data)
	if err != nil {
		fmt.Println("发送数据失败")
		return
	}
	fmt.Printf("一共发送了%d个字节的数据\n", n)
}

// buf := bytes.Buffer{}
// encoder := gob.NewEncoder(&buf)

// err := encoder.Encode(p)
// if err != nil {
// 	fmt.Println("编码失败,错误原因: ", err)
// 	return
// }
