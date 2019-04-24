package main

import (
	"GoServer/Common/config"
	"GoServer/Common/public"
	"GoServer/Common/srpc"
	"GoServer/RedisServer/rpcapi"
	pb "GoServer/Share/Proto/redisrpc"
	"errors"
	"os"

	"github.com/Unknwon/goconfig"
	log "github.com/cihub/seelog"
)

func main() {
	defer public.PanicHandler()
	//load log config file
	_, e := os.Stat("conf/seelog.xml")
	if e != nil {
		log.Error("stat seelog.xml err %v", e)
		return
	}

	logger, err := log.LoggerFromConfigAsFile("conf/seelog.xml")
	if err != nil {
		panic(err)
	}
	log.ReplaceLogger(logger)

	// load config file
	cfg, err := goconfig.LoadConfigFile("./conf/conf.ini")
	if err != nil {
		log.Error("read config error!", "err", err)
		panic(err)
	}

	rpclisaddr, err := cfg.GetValue("", "redis_rpc_listen")
	if err != nil {
		log.Error("read config redis_rpc_listen error!", "err", err)
		panic(err)
	} else {
		log.Info("read config redis_rpc_listen ", rpclisaddr)
	}

	config := &config.Config{
		RpcListen: rpclisaddr,
	}

	// start rpc server
	err = rpcServer(config)
	if err != nil {
		log.Error("rpc server start error ", err)
		panic(err)
	}

	log.Info("redis server start succsessful")

	select {}
}

func rpcServer(config *config.Config) error {
	handle := srpc.NewSrpcServer()
	if handle == nil {
		log.Error("cant new srpc server ")
		return errors.New("cant new srpc server")
	}

	err := handle.StartRedisServer(config.RpcListen, func() error {
		log.Info("register redis server rpc")
		pb.RegisterRedisServer(handle.GrpcServer, &rpcapi.RedisRpcServer{})
		return nil
	})

	if err != nil {
		log.Error("error start redis rpc : ", err)
		return err
	}
	go handle.Serve()
	return nil
}
