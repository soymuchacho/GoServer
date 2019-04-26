package main

import (
	"context"
	"fmt"

	"GoServer/Common/srpc"
	pb "GoServer/Share/Proto/redisrpc"
)

const (
	address = "localhost:50000"
)

func main() {
	fmt.Println("Client Rpc")

	handle := srpc.NewSrpcClient()
	err := handle.StartRpcClient(address)
	if err != nil {
		fmt.Println("Start Rpc Cleint Error")
		return
	}

	defer handle.StopRpcClient()

	c := pb.NewRedisClient(handle.Conn)

	r, err := c.TestRedis(context.Background(), &pb.TestRedisRequest{Test: " Redis Test "})
	if err != nil {
		fmt.Println("Request Redis Test Error ", err)
	}
	fmt.Println("Reply : ", r.Reply)
	r, err = c.TestRedis(context.Background(), &pb.TestRedisRequest{Test: " Redis Test "})
	if err != nil {
		fmt.Println("Request Redis Test Error ", err)
	}
	fmt.Println("Reply : ", r.Reply)
	r, err = c.TestRedis(context.Background(), &pb.TestRedisRequest{Test: " Redis Test "})
	if err != nil {
		fmt.Println("Request Redis Test Error ", err)
	}
	fmt.Println("Reply : ", r.Reply)
	select {}
}
