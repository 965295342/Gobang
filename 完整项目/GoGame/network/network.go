package mynetwork

import (
	"GoGame/data"
	"GoGame/def"
	"GoGame/server"
	"log"
	"net"

	"github.com/golang/protobuf/proto"
)

func init() {
	server.Register(def.CREATE_ROOM, OnCreatRoom)
	server.Register(def.ENROLL, OnEnroll)
	server.Register(def.HEART_BEAT, OnHeartBeat)
}
func GetLocalIPV4() (address string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		//取网络地址的网卡的信息
		ipNet, isIpNet := addr.(*net.IPNet)
		//是网卡并且不是本地环回网卡
		if isIpNet && !ipNet.IP.IsLoopback() {
			ipv4 := ipNet.IP.To4()
			//能正常转成ipv4
			if ipv4 != nil {
				return ipv4.String(), nil
			}
		}
	}
	return "", nil
}

func OnCreatRoom(message *def.NormalMessageC2S) {

}

func OnEnroll(message *def.NormalMessageC2S) {
	_, ok := data.PlayerMap[message.STRING]
	newMessage := getNewMessage()
	newMessage.BOOL = true
	newMessage.ID = 3
	if ok {
		newMessage.BOOL = false
		log.Print("玩家已经注册:", message.STRING)
	} else {
		log.Print("玩家注册成功:", message.STRING)
	}
	player := def.Player{}
	player.Name = message.STRING
	player.Room = ""
	data.PlayerMap[message.STRING] = player

	data, err := proto.Marshal(&newMessage)
	if err != nil {
		log.Print(err)
	}

	server.Send(def.ENROLL, data)

}

func OnHeartBeat(message *def.NormalMessageC2S) {
	newMessage := getNewMessage()

	data, err := proto.Marshal(&newMessage)
	if err != nil {
		log.Print(err)
	}

	server.Send(def.HEART_BEAT, data)
}

func getNewMessage() def.NormalMessageS2C {
	newMessage := def.NormalMessageS2C{}
	newMessage.BOOL = true
	newMessage.ID = 1
	newMessage.INT32 = 1
	newMessage.S2C = true
	newMessage.STRING = "message"
	return newMessage
}
