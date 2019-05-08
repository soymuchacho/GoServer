package main

import (
	"network"

	log "github.com/cihub/seelog"
)

var times uint64 = 0

type NetHandler struct {
}

func (h NetHandler) OnRecv(sess *network.Session, msg []byte) {
	//recvP := network.Reader(msg)
	//mtype, _ := recvP.ReadU16()
	//str, _ := recvP.ReadString()
	//log.Debug("recv type ", mtype, " msg : ", str)
	times++
	log.Debug("recv times ", times)
	//writer := network.Writer()
	//network.Pack(1, "world", writer)

	//sess.Send(writer)
	//log.Debug("end")
}

func (h NetHandler) OnConnect(sess *network.Session) {
	log.Debug("client connect : ip [", sess.GetIp(), "] port [", sess.GetPort(), "]")
}

func (h NetHandler) OnDisConnect(sess *network.Session) {
	log.Debug("client disconnect : ip [", sess.GetIp(), "] port [", sess.GetPort(), "]")
}
