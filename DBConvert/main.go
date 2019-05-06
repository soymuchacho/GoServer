package main

import (
	"fmt"
	"os"
	"public"
	"runtime"
	"strings"

	"github.com/Unknwon/goconfig"
	log "github.com/cihub/seelog"
	"github.com/gohouse/converter"
)

func main() {
	defer public.PanicHandler()
	defer log.Flush()
	//load log config file
	_, e := os.Stat("conf/seelog.xml")
	if e != nil {
		log.Error("stat seelog.xml err %v", e)
		return
	}

	logger, err := log.LoggerFromConfigAsFile("conf/seelog.xml")
	if err != nil {
		panic(err)
	}
	log.ReplaceLogger(logger)

	// load config file
	cfg, err := goconfig.LoadConfigFile("./conf/conf.ini")
	if err != nil {
		log.Error("read config error!", "err", err)
		panic(err)
	}

	savepath, err := cfg.GetValue("", "savepath")
	if err != nil {
		log.Error("read config savepath error!", "err", err)
		panic(err)
	} else {
		log.Info("read config savepath ", savepath)
	}

	host, err := cfg.GetValue("", "dbhost")
	if err != nil {
		log.Error("read config dbhost error!", "err", err)
		panic(err)
	} else {
		log.Info("read config dbhost ", host)
	}

	username, err := cfg.GetValue("", "username")
	if err != nil {
		log.Error("read config username error!", "err", err)
		panic(err)
	} else {
		log.Info("read config username ", username)
	}

	password, err := cfg.GetValue("", "password")
	if err != nil {
		log.Error("read config password error!", "err", err)
		panic(err)
	} else {
		log.Info("read config password ", password)
	}

	databases, err := cfg.GetValue("", "databases")
	if err != nil {
		log.Error("read config databases error!", "err", err)
		panic(err)
	} else {
		log.Info("read config databases ", databases)
	}

	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Warn("Cpu number: ", runtime.NumCPU())

	dbs := strings.Split(databases, ",")
	for _, db := range dbs {
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, db)
		log.Debug("dbbase : ", dsn, " db : ", db)

		// initialize
		t2t := converter.NewTable2Struct()
		// Persionalized configuration
		t2t.Config(&converter.T2tConfig{
			// it is not add tag if the first letter is upper. default false add tag ,true not add.
			RmTagIfUcFirsted: false,
			// it will convert the tag field who has upper letters to lower. default false not convert.
			TagToLower: false,
			// If the first letter of the field is capitalized, convert other letters to lowercase. default false not convert.
			UcFirstOnly: false,
			//// 每个struct放入单独的文件,默认false,放入同一个文件(暂未提供)
			SeperatFile: true,
		})

		// 开始迁移转换
		err = t2t.
			// 指定某个表,如果不指定,则默认全部表都迁移
			Table("").
			// 表前缀
			Prefix("table_").
			// 是否添加json tag
			EnableJsonTag(false).
			// 生成struct的包名(默认为空的话, 则取名为: package model)
			PackageName("dbmodel").
			// tag字段的key值,默认是orm
			TagKey("gorm").
			// whether to add a structure method to get the table name
			RealNameMethod("TableName").
			// the path where to save the generated structure
			SavePath(fmt.Sprintf("%s/%s.go", savepath, db)).
			// mysql dsn,or use the method of t2t.DB(),the param of the method is *sql.DB object
			Dsn(dsn).
			// execute
			Run()

		if err != nil {
			log.Error(err)
		}
	}

}
