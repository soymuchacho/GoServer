package db

import (
	"GoServer/Common/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/cihub/seelog"
)

type Driver struct {
	dbsMap map[string]*DB
}

type MQData struct {
	Key   string
	Value string
}

type CommandData struct {
	DbName string
	Table  string
	Where  string
	Condition string
	Datas  []MQData
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

func (this *Driver) Query(data CommandData) error {
	for db,ok := dbsMap[data.DbName]; ok {
		stmt, err := db.Prepare("select ? from ? where ? ?")
		if err != nil {
			return err
		}

		rows, err := stmt.Exec(data.Datas, data.Table, data.Where, data.Condition)
		if err != nil {
			return err
		}

		for rows.Next() {

		}
	}

	return errors.New("db is not open")
}


func (this *Driver) Insert(data CommandData) error {
	for db,ok := dbsMap[data.DbName]; ok {
		stmt, err := db.Prepare("select ? from ? where ? ?")
		if err != nil {
			return err
		}

		rows, err := stmt.Exec(data.Datas, data.Table, data.Where, data.Condition)
		if err != nil {
			return err
		}

		for rows.Next() {

		}
	}

	return errors.New("db is not open")
}



func (this *Driver) Update(data CommandData) error {
	for db,ok := dbsMap[data.DbName]; ok {
		stmt, err := db.Prepare("select ? from ? where ? ?")
		if err != nil {
			return err
		}

		rows, err := stmt.Exec(data.Datas, data.Table, data.Where, data.Condition)
		if err != nil {
			return err
		}

		for rows.Next() {

		}
	}

	return errors.New("db is not open")
}



func (this *Driver) Delete(data CommandData) error {
	for db,ok := dbsMap[data.DbName]; ok {
		stmt, err := db.Prepare("select ? from ? where ? ?")
		if err != nil {
			return err
		}

		string query := func(datas []MQData) string {
			var res string = " "
			for da := range datas {
				res.append(" ", da.Key, "=", da.Value, " " )
			} 
			return res
		}(data.Datas)

		rows, err := stmt.Exec(data.Datas, data.Table, data.Where, data.Condition)
		if err != nil {
			return err
		}

		for rows.Next() {

		}
	}

	return errors.New("db is not open")
}


