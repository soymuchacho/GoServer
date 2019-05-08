package main

import (
	"config"
	"net/http"
	"network"
	"os"
	"runtime"
	"time"

	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
)

func main() {
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

	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Warn("Cpu number: ", runtime.NumCPU())

	config := &config.Config{
		ServName: "TestServer",
		NetCfg: &config.NetworkCfg{
			Name:         "",
			NetType:      "tcp",
			ConnType:     "connect",
			Address:      "127.0.0.1:50001",
			SockBuf:      32767,
			ReadDeadline: 10 * time.Second,
		},
		HttpCfg: &config.HttpServCfg{
			Name:    "httpserv",
			Address: "127.0.0.1:8000",
		},
	}

	// Test Client
	go func() {
		var handler NetHandler
		netDriver := network.NewNetDriver(handler)
		err := netDriver.TcpConnect(config)
		if err != nil {
			log.Debug("TcpConnect error : ", err)
			return
		}
	}()

	// Test Http
	hpDrive := network.NewHttpDriver()
	hpDrive.AddRouter("GET", "/test", func(c *gin.Context) {
		log.Debug("Get /test")
		c.String(http.StatusOK, "test1 OK")
	})

	hpDrive.HttpServer(config)
	select {}
}
