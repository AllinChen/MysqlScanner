package tablescanner

import (
	"database/sql"
	"fmt"

	"github.com/AllinChen/MysqlScanner/config"
	_ "github.com/go-sql-driver/mysql"
)

type DBInfoTool struct {
}

var DBTool DBInfoTool

// GetDBConnt 建立数据库连接
func (DBInfoTool) GetDBConnt() (*sql.DB, error) {

	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", config.Conf.DB.User, config.Conf.DB.Password, config.Conf.DB.Address, config.Conf.DB.DBName))
	if err != nil {
		return nil, fmt.Errorf("open db conn err: %s", err.Error())
	}

	return db, nil
}
