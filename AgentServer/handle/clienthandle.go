package handle

import (
	"GoServer/Common/network"

	log "github.com/cihub/seelog"
)

type Handler struct {
}

var times uint64 = 0

func (h Handler) OnRecv(sess *network.Session, msg []byte) {
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

func (h Handler) OnConnect(sess *network.Session) {
	log.Debug("client connect : ip [", sess.Ip, "] port [", sess.Port, "]")
}

func (h Handler) OnDisConnect(sess *network.Session) {
	log.Debug("client disconnect : ip [", sess.Ip, "] port [", sess.Port, "]")
}
