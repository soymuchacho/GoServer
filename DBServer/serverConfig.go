package main

import (
	"config"
	"strings"

	"github.com/Unknwon/goconfig"
	log "github.com/cihub/seelog"
)

func init() {
	servCfg = &config.Config{}
}

var gCfg *goconfig.ConfigFile
var servCfg *config.Config

/* = &config.Config{
	RpcCfg: []*config.RpcNetCfg{
		&config.RpcNetCfg{
			Name:     "rpc",
			ConnType: "listen",
			Address:  "127.0.0.1:50002",
		},
	},
	DbCfgs: dbcfg,
}
*/

func serverConfig() (*config.Config, error) {
	// load config file
	cfg, err := goconfig.LoadConfigFile("./conf/conf.ini")
	if err != nil {
		log.Error("read config error!", err)
		return nil, err
	}

	gCfg = cfg
	err = rpcConfig()
	if err != nil {
		return servCfg, err
	}

	err = mysqlConfig()
	if err != nil {
		return servCfg, err
	}

	return servCfg, err
}

func rpcConfig() error {
	var rpccfg []*config.RpcNetCfg
	rpcaddr, err := gCfg.GetValue("rpc", "rpcaddress")
	if err != nil {
		log.Error("no rpc rpcaddress config !")
		return err
	} else {
		log.Debug("read config rpc rpcaddress ", rpcaddr)
	}

	rpccfg = append(rpccfg, &config.RpcNetCfg{
		Name:     "dbrpc",
		ConnType: "listen",
		Address:  rpcaddr,
	})

	servCfg.RpcCfg = rpccfg
	return nil
}

func mysqlConfig() error {
	var dbcfg []*config.DbServCfg

	sqladdr, err := gCfg.GetValue("mysql", "address")
	if err != nil {
		log.Error("no mysql address config !")
		return err
	} else {
		log.Debug("read config mysql address ", sqladdr)
	}

	sqlcharset, err := gCfg.GetValue("mysql", "charset")
	if err != nil {
		log.Error("no mysql charset config !")
		return err
	} else {
		log.Debug("read config mysql charset ", sqlcharset)
	}

	sqluser, err := gCfg.GetValue("mysql", "username")
	if err != nil {
		log.Error("no mysql username config !")
		return err
	} else {
		log.Debug("read config mysql username ", sqluser)
	}

	sqlpwd, err := gCfg.GetValue("mysql", "password")
	if err != nil {
		log.Error("no mysql password config !")
		return err
	} else {
		log.Debug("read config mysql password ", sqlpwd)
	}

	sqldb, err := gCfg.GetValue("mysql", "databases")
	if err != nil {
		log.Error("no mysql databases config !")
		return err
	} else {
		log.Debug("read config mysql databases ", sqldb)
	}

	sqldbs := strings.Split(sqldb, ",")
	for _, dbname := range sqldbs {
		dbcfg = append(dbcfg, &config.DbServCfg{
			DBType:   "mysql",
			Protocol: "tcp",
			Addr:     sqladdr,
			User:     sqluser,
			Pwd:      sqlpwd,
			DataBase: dbname,
			Charset:  "utf8",
		})
	}

	servCfg.DbCfgs = dbcfg
	return nil
}
