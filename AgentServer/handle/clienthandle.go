package handle

import (
	"GoServer/Common/network"
	"encoding/binary"

	log "github.com/cihub/seelog"
)

type Handler struct {
}

func (h Handler) OnRecv(sess *network.Session, msg []byte) {
	recvP := network.Reader(msg)
	mtype, _ := recvP.ReadU16()
	str, _ := recvP.ReadString()
	log.Debug("recv type ", mtype, " msg : ", str)

	writer := network.Writer()
	writer.WriteU16(0)
	pack := network.Pack(1, "helloworld", nil)
	binary.BigEndian.PutUint16(pack[:2], uint16(len(pack)-2))
	sess.Send(pack)
	log.Debug("end")
}

func (h Handler) OnConnect(sess *network.Session) {
	log.Debug("client connect : ip [", sess.Ip, "] port [", sess.Port, "]")
}

func (h Handler) OnDisConnect(sess *network.Session) {
	log.Debug("client disconnect : ip [", sess.Ip, "] port [", sess.Port, "]")
}
