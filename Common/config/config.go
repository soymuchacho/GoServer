package config

type Config struct {
	ServName  string
	RpcListen string
	TcpListen string
	UdpListen string
	TcpConn   string
	UpdConn   string
	Sockbuf   int32
}
