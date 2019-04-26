package network

import (
	"encoding/binary"
	"io"
	"net"
	"time"

	"GoServer/Common/config"

	log "github.com/cihub/seelog"
)

func TcpServer(config *config.Config, ioop IOOperate) error {
	// resolve address & start listening
	tcpAddr, err := net.ResolveTCPAddr("tcp4", config.TcpListen)
	if err != nil {
		return err
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return err
	}

	log.Debug("listening on:", listener.Addr())

	// loop accepting
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Debug("accept failed:", err)
			continue
		}
		// set socket read buffer
		conn.SetReadBuffer(config.SockBuf)
		// set socket write buffer
		conn.SetWriteBuffer(config.SockBuf)

		// start a goroutine for every incoming connection for reading
		go handleClient(conn, config, ioop)
	}
}

func handleClient(conn net.Conn, config *config.Config, ioop IOOperate) error {
	defer conn.Close()

	host, port, err := net.SplitHostPort(conn.RemoteAddr().String())
	if err != nil {
		log.Error("cannot get remote address:", err)
		return err
	}

	sess := &Session{
		Conn:     conn,
		ConnTime: time.Now().Unix(),
		Ip:       string(net.ParseIP(host)),
		Port:     port,
	}

	go sess.Start(config, ioop)

	header := make([]byte, 2)
	// read loop
	for {
		// solve dead link problem:
		// physical disconnection without any communcation between client and server
		// will cause the read to block FOREVER, so a timeout is a rescue.
		conn.SetReadDeadline(time.Now().Add(config.ReadDeadline))

		// read 2B header
		n, err := io.ReadFull(conn, header)
		if err != nil {
			log.Debug("read header failed, ip:%v reason:%v size:%v", sess.Ip, err, n)
			return err
		}
		size := binary.BigEndian.Uint16(header)

		// alloc a byte slice of the size defined in the header for reading data
		payload := make([]byte, size)
		n, err = io.ReadFull(conn, payload)
		if err != nil {
			log.Debug("read payload failed, ip:%v reason:%v size:%v", sess.Ip, err, n)
			return err
		}

		// deliver the data to the input queue
		select {
		case sess.In <- payload: // payload queued
		case <-sess.Die:
			log.Debug("connection closed by logic, flag:%v ip:%v", sess.Flag, sess.Ip)
			return nil
		}
	}

}
