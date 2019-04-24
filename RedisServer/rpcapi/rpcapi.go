package rpcapi

import (
	pb "GoServer/Share/Proto/redisrpc"
	"context"
)

type RedisRpcServer struct{}

func (s *RedisRpcServer) TestRedis(ctx context.Context, in *pb.TestRedisRequest) (*pb.TestRedisResponse, error) {
	return &pb.TestRedisResponse{Reply: "test ok" + in.Test}, nil
}
