package config

import "time"

type Config struct {
	ServName     string
	RpcListen    string
	TcpListen    string
	UdpListen    string
	HttpListen   string
	TcpConn      string
	UpdConn      string
	SockBuf      int
	QueueSize    int32
	ReadDeadline time.Duration
}
