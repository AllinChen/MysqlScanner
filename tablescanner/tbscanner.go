package tablescanner

import (
	"fmt"

	"github.com/AllinChen/MysqlScanner/tinyutils"
)

type tablestruct struct {
	TableCatalog   interface{} `bdb:"TABLE_CATALOG"`
	TableSchema    interface{} `bdb:"TABLE_SCHEMA"`
	TableName      interface{} `bdb:"TABLE_NAME"`
	TableType      interface{} `bdb:"TABLE_TYPE"`
	Engine         interface{} `bdb:"ENGINE"`
	Version        interface{} `bdb:"VERSION"`
	RowFormat      interface{} `bdb:"ROW_FORMAT"`
	TableRows      interface{} `bdb:"TABLE_ROWS"`
	AvgRowLength   interface{} `bdb:"AVG_ROW_LENGTH"`
	DataLength     interface{} `bdb:"DATA_LENGTH"`
	MaxDataLength  interface{} `bdb:"MAX_DATA_LENGTH"`
	IndexLength    interface{} `bdb:"INDEX_LENGTH"`
	DataFree       interface{} `bdb:"DATA_FREE"`
	AutoIncrement  interface{} `bdb:"AUTO_INCREMENT"`
	CreateTime     interface{} `bdb:"CREATE_TIME"`
	UpdateTime     interface{} `bdb:"UPDATE_TIME"`
	CheckTime      interface{} `bdb:"CHECK_TIME"`
	TableCollation interface{} `bdb:"TABLE_COLLATION"`
	CheckSum       interface{} `bdb:"CHECKSUM"`
	CreateOptions  interface{} `bdb:"CREATE_OPTIONS"`
	TableComment   interface{} `bdb:"TABLE_COMMENT"`
}

func GetTables(dbName string) ([]table, error) {
	sql := fmt.Sprintf("select * from information_schema.tables where table_schema = '%s'", dbName)

	db, err := DBTool.GetDBConnt()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tbs []tablestruct
	for rows.Next() {
		var tb tablestruct
		if err = rows.Scan(
			&tb.TableCatalog,
			&tb.TableSchema,
			&tb.TableName,
			&tb.TableType,
			&tb.Engine,
			&tb.Version,
			&tb.RowFormat,
			&tb.TableRows,
			&tb.AvgRowLength,
			&tb.DataLength,
			&tb.MaxDataLength,
			&tb.IndexLength,
			&tb.DataFree,
			&tb.AutoIncrement,
			&tb.CheckTime,
			&tb.UpdateTime,
			&tb.CheckTime,
			&tb.TableCollation,
			&tb.CheckSum,
			&tb.CreateOptions,
			&tb.TableComment,
		); nil != err {
			return nil, err
		}
		tbs = append(tbs, tb)

	}

	return transTables(tbs), nil
}

type table struct {
	TableCatalog   string
	TableSchema    string
	TableName      string
	TableType      string
	Engine         string
	Version        string
	RowFormat      string
	TableRows      string
	AvgRowLength   string
	DataLength     string
	MaxDataLength  string
	IndexLength    string
	DataFree       string
	AutoIncrement  string
	CreateTime     string
	UpdateTime     string
	CheckTime      string
	TableCollation string
	CheckSum       string
	CreateOptions  string
	TableComment   string
}

func transTables(tbfs []tablestruct) (tbs []table) {
	for _, tbf := range tbfs {
		var tb table
		tb.TableCatalog = tinyutils.TransValue(tbf.TableCatalog)
		tb.TableSchema = tinyutils.TransValue(tbf.TableSchema)
		tb.TableName = tinyutils.TransValue(tbf.TableName)
		tb.TableType = tinyutils.TransValue(tbf.TableType)
		tb.Engine = tinyutils.TransValue(tbf.Engine)
		tb.Version = tinyutils.TransValue(tbf.Version)
		tb.RowFormat = tinyutils.TransValue(tbf.RowFormat)
		tb.TableRows = tinyutils.TransValue(tbf.TableRows)
		tb.AvgRowLength = tinyutils.TransValue(tbf.AvgRowLength)
		tb.DataLength = tinyutils.TransValue(tbf.DataLength)
		tb.MaxDataLength = tinyutils.TransValue(tbf.MaxDataLength)
		tb.IndexLength = tinyutils.TransValue(tbf.IndexLength)
		tb.DataFree = tinyutils.TransValue(tbf.DataFree)
		tb.AutoIncrement = tinyutils.TransValue(tbf.AutoIncrement)
		tb.CreateTime = tinyutils.TransValue(tbf.CreateTime)
		tb.UpdateTime = tinyutils.TransValue(tbf.UpdateTime)
		tb.CheckTime = tinyutils.TransValue(tbf.CheckTime)
		tb.TableCollation = tinyutils.TransValue(tbf.TableCollation)
		tb.CheckSum = tinyutils.TransValue(tbf.CheckSum)
		tb.CreateOptions = tinyutils.TransValue(tbf.CreateOptions)
		tb.TableComment = tinyutils.TransValue(tbf.TableComment)
		tbs = append(tbs, tb)
	}
	return
}
