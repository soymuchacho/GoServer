package db

import (
	"GoServer/Common/config"
	"database/sql"
	"errors"
	"fmt"

	log "github.com/cihub/seelog"
	_ "github.com/go-sql-driver/mysql"
)

type Driver struct {
	dbsMap map[string]*sql.DB
}

type MQData struct {
	Key   string
	Value string
}

type MQCmd struct {
	DbName    string
	Table     string
	Where     string
	Condition string
	Datas     []MQData
}

type MQResult struct {
	Res string
}

func NewDBDriver() *Driver {
	return &Driver{
		dbsMap: make(map[string]*sql.DB),
	}
}

func (this *Driver) Default(config *config.Config) error {

	for _, cfg := range config.DbCfgs {
		db, err := sql.Open(cfg.DBType, fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8", cfg.User, cfg.Pwd, cfg.Protocol,
			cfg.Addr, cfg.DataBase, cfg.Charset))
		if err != nil {
			log.Error("connect db ", cfg.DataBase, " err : ", err)
			return err
		}

		this.dbsMap[cfg.DataBase] = db
	}
	return nil
}

func (this *Driver) Query(cmd MQCmd) error {
	if db, ok := this.dbsMap[cmd.DbName]; ok {
		stmt, err := db.Prepare("select ? from ? where ? ?")
		if err != nil {
			return err
		}

		rows, err := stmt.Query(cmd.Datas, cmd.Table, cmd.Where, cmd.Condition)
		if err != nil {
			return err
		}

		for rows.Next() {

		}

	}

	return errors.New("db is not open")
}

func (this *Driver) Insert(cmd MQCmd) error {
	/*
		if db, ok := this.dbsMap[cmd.DbName]; ok {
			stmt, err := db.Prepare("select ? from ? where ? ?")
			if err != nil {
				return err
			}
		}
	*/
	return errors.New("db is not open")
}

func (this *Driver) Update(cmd MQCmd) error {
	/*
		if db, ok := this.dbsMap[cmd.DbName]; ok {
			stmt, err := db.Prepare("select ? from ? where ? ?")
			if err != nil {
				return err
			}
		}
	*/

	return errors.New("db is not open")
}

func (this *Driver) Delete(cmd MQCmd) error {
	/*
		if db, ok := this.dbsMap[cmd.DbName]; ok {
			stmt, err := db.Prepare("select ? from ? where ? ?")
			if err != nil {
				return err
			}

		}
	*/

	return errors.New("db is not open")
}
