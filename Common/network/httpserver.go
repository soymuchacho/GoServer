package network

import (
	"GoServer/Common/config"

	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
)

type HttpConfig struct {
	routersMp map[string]map[string]gin.HandlerFunc
}

func NewHttpConfig() *HttpConfig {
	return &HttpConfig{
		routersMp: make(map[string]map[string]gin.HandlerFunc),
	}
}

func (this *HttpConfig) AddRouter(hType string, hPath string, cb gin.HandlerFunc) {
	if _, ok := this.routersMp[hType]; ok {
		this.routersMp[hType][hPath] = cb
	} else {
		this.routersMp[hType] = make(map[string]gin.HandlerFunc)
		this.routersMp[hType][hPath] = cb
	}
}

func HttpServer(config *config.Config, hpConfig *HttpConfig) error {
	router := gin.Default()

	for httpType, value := range hpConfig.routersMp {
		switch httpType {
		case "GET":
			for httpPath, funcs := range value {
				router.GET(httpPath, funcs)
			}
		case "POST":
			for httpPath, funcs := range value {
				router.POST(httpPath, funcs)
			}
		default:
			log.Error("cant add this router : ", httpType)
		}
	}

	go router.Run(config.HttpListen)
	return nil
}
