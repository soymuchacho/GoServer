package handle

import (
	"GoServer/Common/network"

	log "github.com/cihub/seelog"
)

type Handler struct {
}

func (h Handler) OnRecv(sess *network.Session, msg []byte) {
	log.Debug("Recv msg size ", len(msg))
}

func (h Handler) OnConnect(sess *network.Session) {
	log.Debug("OnConnect")
}

func (h Handler) OnDisConnect(sess *network.Session) {
	log.Debug("OnDisConnect")
}
