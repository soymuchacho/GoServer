package main

import (
	"os"
	"public"
	"runtime"
	"time"

	"config"
	"network"

	"github.com/Unknwon/goconfig"
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

	// load config file
	cfg, err := goconfig.LoadConfigFile("./conf/conf.ini")
	if err != nil {
		log.Error("read config error!", "err", err)
		panic(err)
	}

	tcplisaddr, err := cfg.GetValue("tcp", "agentListen")
	if err != nil {
		log.Error("read config agentListen error!", "err", err)
		panic(err)
	} else {
		log.Info("read config agentListen ", tcplisaddr)
	}

	dbrpcaddr, err := cfg.GetValue("rpc", "dbservice")
	if err != nil {
		log.Error("read config rpc dbservice error!", "err", err)
		panic(err)
	} else {
		log.Info("read config rpc dbservice ", dbrpcaddr)
	}

	var rpcinfo []*config.RpcNetCfg
	rpcconfig := &config.RpcNetCfg{
		Name:     "agentdb",
		ConnType: "connect",
		Address:  dbrpcaddr,
	}
	rpcinfo = append(rpcinfo, rpcconfig)

	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Warn("Cpu number: ", runtime.NumCPU())

	config := &config.Config{
		ServName: "AgentServer",
		NetCfg: &config.NetworkCfg{
			Name:         "Agent",
			NetType:      "tcp",
			ConnType:     "listen",
			Address:      tcplisaddr,
			SockBuf:      32767,
			QueueSize:    32767,
			ReadDeadline: 10 * time.Second,
		},
		RpcCfg: rpcinfo,
	}

	var handler NetHandler
	netDriver := network.NewNetDriver(handler)

	go netDriver.TcpServer(config)

	err = InitRpc(config)
	if err != nil {
		log.Error("Init the RPC connection error : ", err)
	}
	select {}
}

func InitRpc(config *config.Config) error {
	log.Debug("Init RPC len ", len(config.RpcCfg))
	for _, rpcinfo := range config.RpcCfg {
		if rpcinfo.ConnType == "connect" {
			// connect
			log.Debug("rpc connect")
			go ConnectRpc(rpcinfo.Name, rpcinfo.Address)
		} else {
			// listen
		}
	}
	return nil
}
