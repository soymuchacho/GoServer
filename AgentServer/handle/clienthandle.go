package handle

import (
	log "github.com/cihub/seelog"
)

type ClientHandler interface {
	RecvHandler(msg []byte) error
}

func RecvHandler(msg []byte) error {
	log.Debug("Recv msg size ", len(msg))
	route()
	return nil
}
