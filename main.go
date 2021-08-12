package main

import (
	"fmt"

	"github.com/AllinChen/MysqlScanner/config"
	"github.com/AllinChen/MysqlScanner/tablescanner"
)

func main() {
	config.InitConfig("")
	dbs, err := tablescanner.GetDatabases(tablescanner.MysqlDefaultDatabase)
	if err == nil {
		fmt.Println(dbs)
	}
	if err != nil {
		fmt.Println(err)
	}

	tbs, err := tablescanner.GetTables("dbinjection")
	fmt.Println(tbs, err)
}
