package rpcapi

import (
	pb "GoServer/Share/Proto/dbrpc"
	"context"
)

type DBRpcServer struct{}

func (s *DBRpcServer) TestDBServer(ctx context.Context, in *pb.TestDBRequest) (*pb.TestDBResponse, error) {
	return &pb.TestDBResponse{Reply: "test ok" + in.Test}, nil
}
