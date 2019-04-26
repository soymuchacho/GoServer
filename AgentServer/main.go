package main

import (
	"GoServer/AgentServer/handle"
	"GoServer/Common/config"
	"GoServer/Common/network"
	"os"
	"public"

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

	tcplisaddr, err := cfg.GetValue("", "agent_tcp_listen")
	if err != nil {
		log.Error("read config agent_tcp_listen error!", "err", err)
		panic(err)
	} else {
		log.Info("read config agent_tcp_listen ", tcplisaddr)
	}

	config := &config.Config{
		TcpListen: tcplisaddr,
	}

	var ioop handle.Handler
	err = network.TcpServer(config, ioop)
	if err != nil {
		log.Error("tcp listen error ", err)
		panic(err)
	}

	select {}
}
