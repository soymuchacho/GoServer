package network

import (
	"bytes"
	"encoding/binary"
	"net"
	"sync"
	"time"

	"GoServer/Common/config"

	log "github.com/cihub/seelog"
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
	outWsig      chan int // session die write
	outRsig      chan int // session die recv
	Ioop         IOOperate
}

func (this *Session) Start(config *config.Config) {
	func() {
		this.SeMutex.Lock()
		defer this.SeMutex.Unlock()

		this.In = make(chan []byte, config.QueueSize)
		this.Out = make(chan []byte)
		this.Die = make(chan int)
		this.outWsig = make(chan int)
		this.outRsig = make(chan int)

		this.ConnTime = time.Now().Unix()
		this.Flag = SESSION_FLAG_CONNECTED
	}()

	go this.connWrite()
	go this.connRecv()

	this.Ioop.OnConnect(this)
}

func (this *Session) connRecv() {
	defer log.Debug("connRecv funtion end")
	for {
		select {
		case msg, ok := <-this.In:
			if ok {
				this.Ioop.OnRecv(this, msg)
			}
		case <-this.outRsig:
			return
		}
	}
}

func (this *Session) connWrite() {
	defer log.Debug("connWrite funtion end")
	for {
		select {
		case msg, ok := <-this.Out:
			if ok {
				_, err := this.Conn.Write(msg)
				if err != nil {
					log.Error("seesion write msg error : ", err)
				} else {
				}
			}
		case <-this.outWsig:
			log.Debug("session write loop end")
			return
		}
	}

}

func (this *Session) Send(pk *Packet) error {
	// add packet head
	var psize uint16 = uint16(pk.Length())

	log.Debug("send packet length ", pk.Length(), " ", psize)
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, psize)
	buf.Write(pk.Data()[:psize])

	this.Out <- buf.Bytes()
	return nil
}

func (this *Session) Close() error {
	this.Ioop.OnDisConnect(this)
	this.outRsig <- 1
	this.outWsig <- 1

	this.Conn.Close()
	log.Warn("session disconnect ip[", this.Ip, "] port[", this.Port, "]")
	return nil
}
