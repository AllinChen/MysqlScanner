package main

import (
	"fmt"

	"github.com/AllinChen/MysqlScanner/config"
	"github.com/AllinChen/MysqlScanner/markdown"
	"github.com/AllinChen/MysqlScanner/tablescanner"
)

func main() {
	config.InitConfig("")
	StartScanner()

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
		path += config.Conf.DB.Address + ".md"
	}
	fmt.Println(markdown.SaveAsMdFile(path, content))

	return true
}
