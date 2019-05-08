package main

import (
	"GoServer/Common/network"
	"time"

	log "github.com/cihub/seelog"
)

type NetHandler struct {
}

func (h NetHandler) OnRecv(sess *network.Session, msg []byte) {
	recvP := network.Reader(msg)
	mtype, _ := recvP.ReadU16()
	str, _ := recvP.ReadString()
	log.Debug("recv type ", mtype, " msg : ", str)
}

func (h NetHandler) OnConnect(sess *network.Session) {
	log.Debug("connect to server")

	go func() {
		for {
			ticker := time.NewTicker(1 * time.Second)
			select {
			case <-ticker.C:
				for i := 0; i < 100000; i++ {
					writer := network.Writer()
					sendmsg := make([]byte, 65000)
					network.Pack(1, sendmsg, writer)
					//log.Debug("packet length ", writer.Length())
					sess.Send(writer)
				}
			}
		}
	}()
}

func (h NetHandler) OnDisConnect(sess *network.Session) {
	log.Debug("OnDisConnect")
}
