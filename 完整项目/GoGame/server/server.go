package server

import (
	"GoGame/def"
	"fmt"
	"net"
	"sync"
	"time"
)

var CallBackMap map[int32]func(*def.Message)

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
