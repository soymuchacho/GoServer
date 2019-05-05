package main

import (
	"GoServer/DBServer/modle"
	"fmt"
)

func main() {
	modle.ConvertTables()
	fmt.Println("this is the dbserver")
}
