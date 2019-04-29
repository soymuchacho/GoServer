package db

import (
	"GoServer/Common/config"
	"database/sql"
	"fmt"
)

type Driver struct {
	dbsMap map[string]*DB
}

func NewDBDriver() *Driver {
	return &Driver{
		dbsMap: make(map[string]*DB),
	}
}

func (this *Driver) Default(config *config.Config) error {

	db, err := sql.Open(config.DbCfg.DBType, fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8", config.DbCfg.User, config.DbCfg.Pwd, config.DbCfg.Protocol,
		config.DbCfg.Addr, config.DbCfg.DataBase, config.DbCfg.Charset))
	if err != nil {
		log.Error(err)
		return err
	}

}
