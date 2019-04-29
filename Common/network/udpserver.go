package network

import (
	"GoServer/Common/config"

	log "github.com/cihub/seelog"
	"github.com/xtaci/kcp-go"
)

func (this *NetDriver) UdpServer(config *config.Config) error {
	l, err := kcp.Listen(config.NetCfg.Address)
	if err != nil {
		return err
	}

	log.Info("udp listening on:", l.Addr())
	lis := l.(*kcp.Listener)

	if err := lis.SetReadBuffer(config.NetCfg.SockBuf); err != nil {
		log.Debug("SetReadBuffer", err)
	}
	if err := lis.SetWriteBuffer(config.NetCfg.SockBuf); err != nil {
		log.Debug("SetWriteBuffer", err)
	}
	if err := lis.SetDSCP(config.NetCfg.Dscp); err != nil {
		log.Debug("SetDSCP", err)
	}

	// loop accepting
	for {
		conn, err := lis.AcceptKCP()
		if err != nil {
			log.Warn("accept failed:", err)
			continue
		}
		// set kcp parameters
		conn.SetWindowSize(config.NetCfg.SndWnd, config.NetCfg.RcvWnd)
		conn.SetNoDelay(config.NetCfg.Nodelay, config.NetCfg.Interval, config.NetCfg.Resend, config.NetCfg.Nc)
		conn.SetStreamMode(true)
		conn.SetMtu(config.NetCfg.Mtu)

		// start a goroutine for every incoming connection for reading
		go this.HandleClient(conn, config)
	}
}
