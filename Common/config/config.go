package config

import "time"

type DbServCfg struct {
	DBType   string
	Protocol string
	Addr     string
	User     string
	Pwd      string
	DataBase string
	Charset  string
}

type NetworkCfg struct {
	Name         string
	NetType      string // "tcp" "udp" "rpc" "http"
	ConnType     string // "listen" "connect"
	Address      string
	SockBuf      int
	QueueSize    int
	ReadDeadline time.Duration

	// udp use
	Dscp     int
	SndWnd   int
	RcvWnd   int
	Nodelay  int
	Interval int
	Resend   int
	Nc       int
	Mtu      int
}

type HttpServCfg struct {
	Name    string
	Address string
}

type Config struct {
	ServName string
	NetCfg   *NetworkCfg
	DbCfgs   []*DbServCfg // one DBServCfg corresponds to one db connection
	HttpCfg  *HttpServCfg
}
