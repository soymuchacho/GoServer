package main

import (
	"config"
	"db"
	"errors"
	"os"
	"public"
	"srpc"
	"time"

	pb "dbrpc"

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

	dbdriver.GetDB("test").AutoMigrate(&User{})

	var user User
	dbdriver.GetDB("test").First(&user, 1) // find product with id 1
	log.Debug("name ", user.Name, " id ", user.ID)
	dbdriver.GetDB("test").First(&user, "name = ?", "mytest") // find product with code l1212
	log.Debug("name ", user.Name, " id ", user.ID)
	return nil
}

func rpcServer(config *config.Config) error {
	log.Debug("rpc server len ", len(config.RpcCfg))
	for _, cfg := range config.RpcCfg {

		if cfg.ConnType == "listen" {
			log.Debug("listen ", cfg.Address)
			handle := srpc.NewSrpcServer()
			if handle == nil {
				log.Error("cant new srpc server ")
				return errors.New("cant new srpc server")
			}

			err := handle.StartRedisServer(cfg.Address, func() error {
				log.Info("register db server rpc")
				pb.RegisterDBServer(handle.GrpcServer, &DBRpcServer{})
				return nil
			})

			if err != nil {
				log.Error("error start db rpc : ", err)
				return err
			}

			go handle.Serve()
			log.Debug("rpc start : ", cfg.Address)
		} else {
			log.Debug("connect ", cfg.Address)
		}
	}
	return nil
}
