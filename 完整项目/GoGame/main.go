package main

import (
	mynetwork "GoGame/network"
	"GoGame/server"
	"log"
)

func main() {
	myaddress, err := mynetwork.GetLocalIPV4()
	if err != nil {
		log.Println(err)
	}
	log.Println(myaddress)

	_, err = server.NewSocketService()
	if err != nil {
		log.Println(err)
	}

}
