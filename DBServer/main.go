package main

import (
	"GoServer/Common/config"
	"GoServer/Common/db"
	"GoServer/Common/srpc"
	"GoServer/DBServer/model"
	"GoServer/DBServer/rpcapi"
	"errors"
	"os"
	"public"
	"time"

	pb "GoServer/Share/Proto/dbrpc"

	log "github.com/cihub/seelog"
)

func main() {
	defer public.PanicHandler()
	defer log.Flush()
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

	retcfg, err := serverConfig()
	if err != nil {
		log.Error("HandleConfig error!", err)
		panic(err)
	}

	// rpc service
	rpcServer(retcfg)

	// db service
	dbServer(retcfg)

	log.Debug("DBServer Start")
	for {
		time.Sleep(5 * time.Second)
	}
}

func dbServer(config *config.Config) error {
	dbdriver := db.NewDBDriver()
	err := dbdriver.Default(config)
	if err != nil {
		log.Error(err)
		return err
	}

	dbdriver.GetDB("test").AutoMigrate(&model.User{})

	var user model.User
	dbdriver.GetDB("test").First(&user, 1) // find product with id 1
	log.Debug("name ", user.Name, " id ", user.ID)
	dbdriver.GetDB("test").First(&user, "name = ?", "mytest") // find product with code l1212
	log.Debug("name ", user.Name, " id ", user.ID)
	return nil
}

func rpcServer(config *config.Config) error {
	for _, cfg := range config.RpcCfg {
		if cfg.ConnType == "listen" {
			handle := srpc.NewSrpcServer()
			if handle == nil {
				log.Error("cant new srpc server ")
				return errors.New("cant new srpc server")
			}

			err := handle.StartRedisServer(cfg.Address, func() error {
				log.Info("register redis server rpc")
				pb.RegisterMysqlServer(handle.GrpcServer, &rpcapi.DBRpcServer{})
				return nil
			})

			if err != nil {
				log.Error("error start redis rpc : ", err)
				return err
			}
			go handle.Serve()
		} else {

		}
	}
	return nil
}
