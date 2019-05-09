package network

import (
	"config"
	"encoding/binary"
	"io"
	"net"
	"time"

	log "github.com/cihub/seelog"
)

func (this *NetDriver) HandleClient(conn net.Conn, config *config.Config) error {
	defer conn.Close()

	host, port, err := net.SplitHostPort(conn.RemoteAddr().String())
	if err != nil {
		log.Error("cannot get remote address:", err)
		return err
	}

	sess := NewSession(conn, host, port, config.NetCfg.QueueSize)

	err = func() error {
		defer this.OnDisConnect(sess)
		defer sess.Close()

		go sess.Start(this)

		header := make([]byte, 2)

		this.OnConnect(sess)
		// read loop
		for {
			// solve dead link problem:
			// physical disconnection without any communcation between client and server
			// will cause the read to block FOREVER, so a timeout is a rescue.
			conn.SetReadDeadline(time.Now().Add(config.NetCfg.ReadDeadline))

			// read 2B header
			n, err := io.ReadFull(conn, header)
			if err != nil {
				log.Debug("read header failed Ip ", host, " err ", err, " size ", n)
				return err
			}
			size := binary.BigEndian.Uint16(header)
			//log.Debug("recv msg size ", size)
			// alloc a byte slice of the size defined in the header for reading data
			payload := make([]byte, size)
			n, err = io.ReadFull(conn, payload)
			if err != nil {
				log.Debug("read payload failed, ip: ", host, " reason:", err, " size:", n)
				return err
			}

			// deliver the data to the input queue
			select {
			case sess.in <- payload: // payload queued
			case <-sess.GetDie():
				log.Warn("recv die signal , close this session")
				this.OnDisConnect(sess)
				return nil
			}
		}
		return nil
	}()

	log.Debug("handleClient funtion end")
	return err
}
