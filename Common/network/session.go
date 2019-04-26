package network

import (
	"errors"
	"fmt"
	"net"

	"GoServer/Common/config"
)

type IOOperate interface {
	RecvMsg(msg []byte) error
}

type Session struct {
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

	this.In = make(chan []byte, config.QueueSize)
	this.Out = make(chan []byte)
	this.Die = make(chan int)

	for {
		select {
		case msg, ok := <-this.In:
			fmt.Println("write msg len : ", len(msg))
			if ok {
				ioop.RecvMsg(msg)
			}
		case msg, ok := <-this.Out:
			fmt.Println("write msg len : ", len(msg))
			if ok {
				n, err := this.Conn.Write(msg)
				if err != nil {
					fmt.Println("connection write error : ", err)
				} else {
					fmt.Println("connection write byte size ", n)
				}
			}
		case <-this.Die:
			return
		}
	}
}

func (this *Session) Send(msgtype int16, pb interface{}) error {
	msg := Pack(msgtype, pb, nil)
	if msg == nil {
		fmt.Println("send error pack is nil")
		return errors.New("send error : pack is nil")
	}
	this.Out <- msg
	return nil
}
