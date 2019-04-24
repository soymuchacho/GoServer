package srpc

import (
	"errors"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type SrpcServer struct {
	GrpcServer *grpc.Server
	GrpcMut    sync.Mutex
	Lis        net.Listener
}

func init() {

}

func NewSrpcServer() *SrpcServer {
	return &SrpcServer{}
}

func (this *SrpcServer) StartRedisServer(address string, registerCB func() error) error {
	dblStartError := func() error {
		this.GrpcMut.Lock()
		defer this.GrpcMut.Unlock()
		if this.GrpcServer != nil {
			return errors.New("Rpc Server Already Started")
		}
		var err error
		this.Lis, err = net.Listen("tcp", address)
		if err != nil {
			return err
		}
		this.GrpcServer = grpc.NewServer()
		return nil
	}()

	if dblStartError != nil {
		return dblStartError
	}

	err := registerCB()
	return err
}

func (this *SrpcServer) Serve() error {
	err := this.GrpcServer.Serve(this.Lis)
	return err
}

func (this *SrpcServer) StopRedisServer() {
	tmpServer := func() *grpc.Server {
		this.GrpcMut.Lock()
		defer this.GrpcMut.Unlock()
		svr := this.GrpcServer
		if this.GrpcServer != nil {
			this.GrpcServer = nil
		}
		return svr
	}()
	if tmpServer == nil {
		return
	}
	tmpServer.Stop()
}
