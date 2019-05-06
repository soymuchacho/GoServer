package db

import (
	"GoServer/Common/config"
	"fmt"
	"reflect"

	log "github.com/cihub/seelog"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Driver struct {
	dbsMap  map[string]*gorm.DB
	comment chan MQComment
}

type Element struct {
	Key   string
	Value interface{}
}

type Fields struct {
	Fs []*Element
}

type MQComment struct {
	cmttype string // "update", "delete", "insert","select"
	dbName  string
	table   string
	sql     string
	dv      *Driver
}

type MQResult struct {
	rets []*Fields
}

func FieldEncode(key string, val interface{}) (ret string) {
	ty := reflect.TypeOf(val)
	switch ty.Kind() {
	case reflect.String:
		sql := ty.String()
		ret = key + "=" + "'" + sql + "'"
	default:
		ret = fmt.Sprintf("%s=%v", key, val)
	}
	return ret
}

func (dv *Driver) CreateComment(cmttype string, db string, table string, f *Fields) *MQComment {
	var sql_main string
	var sql string

	cmd := &MQComment{
		cmttype: cmttype, // "update", "delete", "insert","select"
		dbName:  db,
		table:   table,
		dv:      dv,
	}

	switch cmttype {
	case "select":
		sql = "select "
		if f == nil || len(f.Fs) == 0 {
			sql_main = "*"
		} else {
			for key, _ := range f.Fs {
				sql_main += fmt.Sprintf("%s", key)
			}
		}
		sql += sql_main
		sql += " from " + table + " "
		cmd.sql = sql
	case "update":
	case "delete":
	case "insert":
	}

	return cmd
}

func (this *MQComment) Where(sql string) *MQComment {
	this.sql += " where " + sql
	return this
	/*
		if f == nil || len(f.Fs) == 0 {
			log.Debug("f is nil or len f.Fs == 0 ", len(f.Fs))
			return this
		} else {
			var count int = 0
			var sql_where string = " where "
			for _, elem := range f.Fs {
				count++
				sql_where += FieldEncode(elem.Key, elem.Value)
				if len(f.Fs) > count {
					sql_where += " and "
				}
			}
			this.sql += sql_where + " "
		}
		return this
	*/
}

func NewDBDriver() *Driver {
	return &Driver{
		dbsMap:  make(map[string]*gorm.DB),
		comment: make(chan MQComment),
	}
}

func (this *Driver) Default(config *config.Config) error {
	for _, cfg := range config.DbCfgs {
		dsn := fmt.Sprintf("%s:%s@%s(%s)/%s?charset=%s&parseTime=True&loc=Local", cfg.User, cfg.Pwd, cfg.Protocol,
			cfg.Addr, cfg.DataBase, cfg.Charset)
		db, err := gorm.Open(cfg.DBType, dsn)
		if err != nil {
			log.Error("connect db ", cfg.DataBase, " err : ", err)
			return err
		}

		log.Debug("open db ok : ", dsn)
		this.dbsMap[cfg.DataBase] = db
	}

	return nil
}

func (this *Driver) Close() {
	for _, db := range this.dbsMap {
		db.Close()
	}
}

func (this *Driver) GetDB(dbname string) *gorm.DB {
	if db, ok := this.dbsMap[dbname]; ok {
		return db
	}
	return nil
}
