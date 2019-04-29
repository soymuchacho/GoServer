package network

import (
	"bytes"
	"encoding/binary"
	"net"
	"sync"
	"time"

	log "github.com/cihub/seelog"
)

type NetAPI interface {
	OnConnect(sess *Session)
	OnRecv(sess *Session, msg []byte)
	OnDisConnect(sess *Session)
}

type NetApiBase struct {
}

type NetDriver struct {
	netApi NetAPI
}

func (this *NetDriver) OnConnect(sess *Session) {
	if this.netApi == nil {
		log.Debug("netdriver connect function, please drive it")
	} else {
		this.netApi.OnConnect(sess)
	}
}

func (this *NetDriver) OnRecv(sess *Session, msg []byte) {
	if this.netApi == nil {
		log.Debug("netdriver recv function, please drive it")
	} else {
		this.netApi.OnRecv(sess, msg)
	}
}

func (this *NetDriver) OnDisConnect(sess *Session) {
	if this.netApi == nil {
		log.Debug("netdriver disconnect function, please drive it")
	} else {
		this.netApi.OnDisConnect(sess)
	}
}

func NewNetDriver(api NetAPI) *NetDriver {
	return &NetDriver{
		netApi: api,
	}
}

const (
	SESSION_FLAG_CONNECTED = iota
	SESSION_FLAG_DISCONNED
)

type Session struct {
	seMux   sync.Mutex
	ip      string // session Ip address
	port    string // session Port
	conn    net.Conn
	in      chan []byte
	out     chan []byte
	outWsig chan int // session die write
	outRsig chan int // session die recv

	flag         int
	connTime     int64    // session connect time
	packTime     int64    // session recv pack time
	lastPackTime int64    // session last recv pack time
	die          chan int // session die chan
}

func NewSession(conn net.Conn, ip string, port string, queueSize int) *Session {
	sess := &Session{}

	sess.conn = conn
	sess.ip = ip
	sess.port = port
	sess.in = make(chan []byte, queueSize)
	sess.out = make(chan []byte)
	sess.die = make(chan int)
	sess.outWsig = make(chan int)
	sess.outRsig = make(chan int)

	sess.connTime = time.Now().Unix()
	sess.flag = SESSION_FLAG_CONNECTED

	return sess
}

func (this *Session) GetDie() chan int {
	return this.die
}

func (this *Session) GetIn() chan []byte {
	return this.in
}

func (this *Session) GetIp() string {
	return this.ip
}

func (this *Session) GetPort() string {
	return this.port
}

func (this *Session) Start(ioop *NetDriver) {
	this.seMux.Lock()
	defer this.seMux.Unlock()

	go this.connWrite(ioop)
	go this.connRecv(ioop)

}

func (this *Session) connRecv(ioop *NetDriver) {
	defer log.Debug("connRecv funtion end")
	for {
		select {
		case msg, ok := <-this.in:
			if ok {
				ioop.OnRecv(this, msg)
			}
		case <-this.outRsig:
			return
		}
	}
}

func (this *Session) connWrite(ioop *NetDriver) {
	defer log.Debug("connWrite funtion end")
	for {
		select {
		case msg, ok := <-this.out:
			if ok {
				_, err := this.conn.Write(msg)
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

	this.out <- buf.Bytes()
	return nil
}

func (this *Session) Close() error {
	this.outRsig <- 1
	this.outWsig <- 1

	this.conn.Close()
	log.Warn("session disconnect ip[", this.ip, "] port[", this.port, "]")
	return nil
}
