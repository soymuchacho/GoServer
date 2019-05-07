package srpc

import (
	"sync"

	"google.golang.org/grpc"
)

// rpc client instance
type SrpcClient struct {
	ServID  string // the id is the server of connected
	Address string
	FunHwd  interface{}      // the handle of request rpc function
	Conn    *grpc.ClientConn // the connection handle of the rpc
	GrpcMut sync.Mutex       // mutex
}

func NewSrpcClient(servid string) *SrpcClient {
	newcli := &SrpcClient{
		ServID: servid,
	}

	if cli, ok := SrpcCliMgr.clis[servid]; ok {
		// the rpc connection is already exsit
		cli.StopRpcClient() // close it
		delete(SrpcCliMgr.clis, servid)
	}
	SrpcCliMgr.clis[servid] = newcli
	return newcli
}

func (this *SrpcClient) StartRpcClient(address string) (err error) {
	this.GrpcMut.Lock()
	defer this.GrpcMut.Unlock()
	this.Address = address

	this.Conn, err = grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		if this.Conn != nil {
			this.Conn.Close()
		}
		return err
	}
	return nil
}

func (this *SrpcClient) ReConnect() error {
	this.StopRpcClient()
	this.StartRpcClient(this.Address)
	return nil
}

func (this *SrpcClient) StopRpcClient() error {
	if this.Conn != nil {
		err := this.Conn.Close()
		return err
	}

	if _, ok := SrpcCliMgr.clis[this.ServID]; ok {
		delete(SrpcCliMgr.clis, this.ServID)
		return nil
	}
	return nil
}

// rpc client instance manamger
type SrpcClientMgr struct {
	clis map[string]*SrpcClient
}

var SrpcCliMgr *SrpcClientMgr

func init() {
	SrpcCliMgr = &SrpcClientMgr{
		clis: make(map[string]*SrpcClient),
	}
}

func (this *SrpcClientMgr) GetRpcClient(servid string) *SrpcClient {
	if cli, ok := this.clis[servid]; ok {
		return cli
	}
	return nil
}
