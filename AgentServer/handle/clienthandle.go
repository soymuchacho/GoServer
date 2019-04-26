package handle

import (
	log "github.com/cihub/seelog"
)

type Handler struct {
}

func (h Handler) RecvMsg(msg []byte) error {
	log.Debug("Recv msg size ", len(msg))
	//route()
	return nil
}
