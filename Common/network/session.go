package network

import (
	"fmt"
	"net"
	"sync"
	"time"

	"GoServer/Common/config"
)

type IOOperate interface {
	OnConnect(sess *Session)
	OnRecv(sess *Session, msg []byte)
	OnDisConnect(sess *Session)
}

const (
	SESSION_FLAG_CONNECTED = iota
	SESSION_FLAG_DISCONNED
)

type Session struct {
	SeMutex      sync.Mutex
	Ip           string // session Ip address
	Port         string // session Port
	Conn         net.Conn
	In           chan []byte
	Out          chan []byte
	Flag         int32
	ConnTime     int64    // session connect time
	PackTime     int64    // session recv pack time
	LastPackTime int64    // session last recv pack time
	Die          chan int // session die chan
}

func (this *Session) Start(config *config.Config, ioop IOOperate) {
	func() {
		this.SeMutex.Lock()
		defer this.SeMutex.Unlock()

		this.In = make(chan []byte, config.QueueSize)
		this.Out = make(chan []byte)
		this.Die = make(chan int)
		this.ConnTime = time.Now().Unix()
		this.Flag = SESSION_FLAG_CONNECTED
	}()

	ioop.OnConnect(this)

	for {
		select {
		case msg, ok := <-this.In:
			if ok {
				ioop.OnRecv(this, msg)
			}
		case msg, ok := <-this.Out:
			if ok {
				n, err := this.Conn.Write(msg)
				if err != nil {
					fmt.Println("connection write error : ", err)
				} else {
					fmt.Println("connection write byte size ", n)
				}
			}
		}
	}
}

func (this *Session) Send(msg []byte) error {
	this.Out <- msg
	return nil
}

func (this *Session) Close() error {
	this.Conn.Close()
	return nil
}
