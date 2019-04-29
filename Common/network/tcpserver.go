package network

import (
	"net"

	"GoServer/Common/config"

	log "github.com/cihub/seelog"
)

func (this *NetDriver) TcpServer(config *config.Config) error {
	// resolve address & start listening
	tcpAddr, err := net.ResolveTCPAddr("tcp4", config.NetCfg.Address)
	if err != nil {
		return err
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return err
	}

	log.Info("Listen to addr : ", tcpAddr)

	// loop accepting
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Debug("accept failed:", err)
			continue
		}
		// set socket read buffer
		conn.SetReadBuffer(config.NetCfg.SockBuf)
		// set socket write buffer
		conn.SetWriteBuffer(config.NetCfg.SockBuf)

		// start a goroutine for every incoming connection for reading
		go this.HandleClient(conn, config)
	}
}
