package main

import (
	"flag"
	"fmt"

	"github.com/AllinChen/MysqlScanner/config"
	"github.com/AllinChen/MysqlScanner/markdown"
	"github.com/AllinChen/MysqlScanner/tablescanner"
)

func main() {
	flag.Parse()
	config.InitConfig("")
	fmt.Printf("Address  : %s \n", config.Conf.DB.Address)
	fmt.Printf("Port     : %d \n", config.Conf.DB.Port)
	fmt.Printf("UserName : %s \n", config.Conf.DB.User)
	fmt.Printf("Password : %s \n", config.Conf.DB.Password)
	fmt.Printf("Database : %s \n", *config.Fdb)
	fmt.Printf("Table    : %s \n", *config.Ftb)
	fmt.Printf("scanning ...........\n")
	StartScanner()
	fmt.Printf("Success!")

}

func StartScanner() bool {
	dbs, err := tablescanner.GetDatabases(tablescanner.MysqlDefaultDatabase)
	if err != nil {
		return false
	}
	content := ""
	for _, db := range dbs {
		dbName := db.Database
		content = markdown.BigTitle(content, dbName)
		tbs, _ := tablescanner.GetTables(dbName)
		content = markdown.DatabaseTable(content, dbName, tbs)
		for _, tb := range tbs {
			content = markdown.SmallTitle(content+"\n", tb.TableName)
			cols, _ := tablescanner.GetColumns(dbName, tb.TableName)
			content = markdown.TableColumn(content, dbName, cols)
		}

	}
	path, err := markdown.GetFilePath()
	if err != nil {
		return false
	}
	if err == nil {
		path += config.Conf.DB.Address
		if *config.Fdb != "" {
			path += " db-" + *config.Fdb
		}
		if *config.Ftb != "" {
			path += " tb-" + *config.Ftb
		}
		path += ".md"
	}

	// fmt.Println(markdown.SaveAsMdFile(path, content))

	return true
}
