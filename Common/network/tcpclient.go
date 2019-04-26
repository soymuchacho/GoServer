package network

import (
	"GoServer/Common/config"
	"encoding/binary"
	"io"
	"net"
	"strings"

	log "github.com/cihub/seelog"
)

func TcpConnect(config *config.Config, ioop IOOperate) error {
	// resolve address & start listening
	tcpAddr, err := net.ResolveTCPAddr("tcp4", config.TcpConn)
	if err != nil {
		log.Error("resolve tcp addr error : ", err)
		return err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Error("connect tcp address[%s] error : ", config.TcpConn)
		return err
	}

	address := strings.Split(config.TcpConn, ":")

	sess := &Session{
		Ip:   address[0],
		Port: address[1],
		Conn: conn,
		Ioop: ioop,
	}

	go sess.Start(config)
	defer func() {
		ioop.OnDisConnect(sess)
		sess.Close()
	}()

	header := make([]byte, 2)
	for {
		// solve dead link problem:
		// physical disconnection without any communcation between client and server
		// will cause the read to block FOREVER, so a timeout is a rescue.
		//conn.SetReadDeadline(time.Now().Add(config.ReadDeadline))

		// read 2B header
		n, err := io.ReadFull(conn, header)
		if err != nil {
			log.Error("read header error ip [", sess.Ip, "] err [", err, "] size [", n, "]")
			return err
		}
		size := binary.BigEndian.Uint16(header)
		log.Debug("recv msg size ", size)
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
			return nil
		}
	}
	return nil
}
