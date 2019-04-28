package main

import (
	"GoServer/Common/config"
	"GoServer/Common/network"
	"GoServer/TestServer/handle"
	"runtime"
	"time"

	log "github.com/cihub/seelog"
)

func main() {
	defer log.Flush()

	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Warn("Cpu number: ", runtime.NumCPU())

	config := &config.Config{
		TcpConn:      "127.0.0.1:50001",
		ReadDeadline: 10 * time.Second,
		SockBuf:      32767,
	}

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

	select {}
}
