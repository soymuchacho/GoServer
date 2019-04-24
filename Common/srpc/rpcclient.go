package srpc

import (
	"sync"

	"google.golang.org/grpc"
)

type SrpcClient struct {
	Conn    *grpc.ClientConn
	GrpcMut sync.Mutex
}

func NewSrpcClient() *SrpcClient {
	return &SrpcClient{}
}

func (this *SrpcClient) StartRpcClient(address string) (err error) {
	this.GrpcMut.Lock()
	defer this.GrpcMut.Unlock()

	this.Conn, err = grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		if this.Conn != nil {
			this.Conn.Close()
		}
		return err
	}
	return nil
}

func (this *SrpcClient) StopRpcClient() error {
	if this.Conn != nil {
		err := this.Conn.Close()
		return err
	}
	return nil
}
