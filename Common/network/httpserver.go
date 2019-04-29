package network

import (
	"GoServer/Common/config"

	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
)

type HttpDriver struct {
	routersMp map[string]map[string]gin.HandlerFunc
}

func NewHttpDriver() *HttpDriver {
	return &HttpDriver{
		routersMp: make(map[string]map[string]gin.HandlerFunc),
	}
}

func (this *HttpDriver) AddRouter(hType string, hPath string, cb gin.HandlerFunc) {
	if _, ok := this.routersMp[hType]; ok {
		this.routersMp[hType][hPath] = cb
	} else {
		this.routersMp[hType] = make(map[string]gin.HandlerFunc)
		this.routersMp[hType][hPath] = cb
	}
}

func (this *HttpDriver) HttpServer(config *config.Config) error {
	router := gin.Default()

	for httpType, value := range this.routersMp {
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

	go router.Run(config.HttpCfg.Address)
	return nil
}
