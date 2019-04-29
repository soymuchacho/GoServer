package network

import (
	"GoServer/Common/config"
	"encoding/binary"
	"io"
	"net"
	"strings"

	log "github.com/cihub/seelog"
)

func (this *NetDriver) TcpConnect(config *config.Config) error {
	// resolve address & start listening
	tcpAddr, err := net.ResolveTCPAddr("tcp4", config.NetCfg.Address)
	if err != nil {
		log.Error("resolve tcp addr error : ", err)
		return err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Error("connect tcp address[%s] error : ", config.NetCfg.Address)
		return err
	}

	address := strings.Split(config.NetCfg.Address, ":")

	sess := NewSession(conn, address[0], address[1], config.NetCfg.QueueSize)

	go sess.Start(this)

	defer func() {
		this.OnDisConnect(sess)
		sess.Close()
	}()

	this.OnConnect(sess)

	header := make([]byte, 2)
	for {
		// read 2B header
		n, err := io.ReadFull(conn, header)
		if err != nil {
			log.Error("read header error ip [", address[0], "] err [", err, "] size [", n, "]")
			return err
		}
		size := binary.BigEndian.Uint16(header)
		log.Debug("recv msg size ", size)
		// alloc a byte slice of the size defined in the header for reading data
		payload := make([]byte, size)
		n, err = io.ReadFull(conn, payload)
		if err != nil {
			log.Debug("read payload failed, ip:", address[0], "reason:", err, "size:", n)
			return err
		}

		// deliver the data to the input queue
		select {
		case sess.in <- payload: // payload queued
		case <-sess.die:
			return nil
		}
	}
	return nil
}
