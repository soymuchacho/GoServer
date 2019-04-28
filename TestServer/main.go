package main

import (
	"GoServer/Common/config"
	"GoServer/Common/network"
	"net/http"
	"runtime"
	"time"

	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
)

func main() {
	defer log.Flush()

	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Warn("Cpu number: ", runtime.NumCPU())

	config := &config.Config{
		TcpConn:      "127.0.0.1:50001",
		ReadDeadline: 10 * time.Second,
		SockBuf:      32767,
		HttpListen:   "127.0.0.1:8080",
	}

	// Test Client
	/*
		for i := 0; i <= 2000; i++ {
			go func() {
				var ioop handle.Handler
				err := network.TcpConnect(config, ioop)
				if err != nil {
					log.Debug("TcpConnect error : ", err)
					return
				}
			}()
		}
	*/

	// Test Http
	hpConfig := network.NewHttpConfig()
	hpConfig.AddRouter("GET", "/test", func(c *gin.Context) {
		log.Debug("Get /test")
		c.String(http.StatusOK, "test1 OK")
	})

	network.HttpServer(config, hpConfig)
	select {}
}
