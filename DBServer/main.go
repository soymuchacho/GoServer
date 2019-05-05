package main

import (
	"fmt"

	"GoServer/Common/config"
	"GoServer/Common/db"
)

func main() {
	dbdriver := db.NewDBDriver()
	config := &config.Config{}
	dbdriver.Default(config)
	fmt.Println("this is the dbserver")
}
