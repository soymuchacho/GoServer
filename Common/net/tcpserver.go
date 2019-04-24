package net

import (
	"net"

	"github.com/olivere/elastic/config"
)

func TcpServer(config *config.Config, handleClientCB func(net.Conn, *config.Config)) error {
	// resolve address & start listening
	tcpAddr, err := net.ResolveTCPAddr("tcp4", config.listen)
	if err != nil {
		return err
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return err
	}

	log.Info("listening on:", listener.Addr())

	// loop accepting
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Warning("accept failed:", err)
			continue
		}
		// set socket read buffer
		conn.SetReadBuffer(config.sockbuf)
		// set socket write buffer
		conn.SetWriteBuffer(config.sockbuf)
		// start a goroutine for every incoming connection for reading
		go handleClientCB(conn, config)
	}
}
