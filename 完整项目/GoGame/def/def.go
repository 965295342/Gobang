package def

import (
	"net"
	"sync"
	"time"
)

type Message struct {
	id       int32
	arrayInt []int32
	words    string
}

// Session struct
type Session struct {
	sID      string
	uID      string
	conn     *net.Conn
	settings map[string]interface{}
}

// SocketService struct
type SocketService struct {
	onMessage    func(*Session, *Message)
	onConnect    func(*Session)
	onDisconnect func(*Session, error)
	Sessions     *sync.Map
	HbInterval   time.Duration
	HbTimeout    time.Duration
	Laddr        string
	Status       int
	Listener     net.Listener
	StopCh       chan error
}

const (
	CREATE_ROOM int32 = 1 //开房
)

type Player struct {
}

const (
	// STUnknown Unknown
	STUnknown = iota
	// STInited Inited
	STInited
	// STRunning Running
	STRunning
	// STStop Stop
	STStop
)
