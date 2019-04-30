package db

import (
	"GoServer/Common/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Driver struct {
	dbsMap map[string]*DB
}

type QueryData struct {
	DbName string
	Table  string
}

func NewDBDriver() *Driver {
	return &Driver{
		dbsMap: make(map[string]*DB),
	}
}

func (this *Driver) Default(config *config.Config) error {

	for cfg := range config.DBCfgs {
		db, err := sql.Open(cfg.DBType, fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8", cfg.User, cfg.Pwd, cfg.Protocol,
			cfg.Addr, cfg.DataBase, cfg.Charset))
		if err != nil {
			log.Error("connect db ", cfg.DataBase, " err : ", err)
			return err
		}

		dbsMap[cfg.DataBase] = db
	}
}

func (this *Driver) QueryRow(datas []QueryData) {
	for dbname, db := range dbsMap {

	}
}
