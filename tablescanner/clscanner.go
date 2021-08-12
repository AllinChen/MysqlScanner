package tablescanner

import (
	"fmt"

	"github.com/AllinChen/MysqlScanner/tinyutils"
)

type columnstruct struct {
	TableCatalog           interface{} `bdb:"TABLE_CATALOG"`
	TableSchema            interface{} `bdb:"TABLE_SCHEMA"`
	TableName              interface{} `bdb:"TABLE_NAME"`
	ColumnName             interface{} `bdb:"Column_Name"`
	OrdinalPosition        interface{} `bdb:"ORDINAL_POSITION"`
	ColumnDefault          interface{} `bdb:"COLUMN_DEFAULT"`
	IsNullAble             interface{} `bdb:"IS_NULLABLE"`
	DataType               interface{} `bdb:"DATA_TYPE"`
	CharacterMaximumLength interface{} `bdb:"CHARACTER_MAXIMUM_LENGTH"`
	CharacterOrtetLength   interface{} `bdb:"CHARACTER_OCTET_LENGTH"`
	NumericPrecision       interface{} `bdb:"NUMERIC_PRECISION"`
	NumericScale           interface{} `bdb:"NUMERIC_SCALE"`
	DataTimePrecision      interface{} `bdb:"DATATIME_PRECISION"`
	CharacterSetName       interface{} `bdb:"CHARACTER_SET_NAME"`
	CollationName          interface{} `bdb:"COLLATION_NAME"`
	ColumnType             interface{} `bdb:"COLUMN_TYPE"`
	ColumnKey              interface{} `bdb:"COLUMN_KEY"`
	Extra                  interface{} `bdb:"EXTRA"`
	Privileges             interface{} `bdb:"PRIVILEGES"`
	ColumnComment          interface{} `bdb:"COLUMN_COMMENT"`
	GenerationExpression   interface{} `bdb:"GENERATION_EXPRESSION"`
}

func GetColumns(dbName, tableName string) ([]column, error) {
	sql := fmt.Sprintf("select * from information_schema.columns where table_schema = '%s' and table_name = '%s'; ", dbName, tableName)
	db, err := DBTool.GetDBConnt()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cols []columnstruct
	for rows.Next() {
		var col columnstruct
		if err = rows.Scan(
			&col.TableCatalog,
			&col.TableSchema,
			&col.TableName,
			&col.ColumnName,
			&col.OrdinalPosition,
			&col.ColumnDefault,
			&col.IsNullAble,
			&col.DataType,
			&col.CharacterMaximumLength,
			&col.CharacterOrtetLength,
			&col.NumericPrecision,
			&col.NumericScale,
			&col.DataTimePrecision,
			&col.CharacterSetName,
			&col.CollationName,
			&col.ColumnType,
			&col.ColumnKey,
			&col.Extra,
			&col.Privileges,
			&col.ColumnComment,
			&col.GenerationExpression,
		); nil != err {
			return nil, err
		}
		cols = append(cols, col)
	}
	return transColumn(cols), nil
}

type column struct {
	TableCatalog           string
	TableSchema            string
	TableName              string
	ColumnName             string
	OrdinalPosition        string
	ColumnDefault          string
	IsNullAble             string
	DataType               string
	CharacterMaximumLength string
	CharacterOrtetLength   string
	NumericPrecision       string
	NumericScale           string
	DataTimePrecision      string
	CharacterSetName       string
	CollationName          string
	ColumnType             string
	ColumnKey              string
	Extra                  string
	Privileges             string
	ColumnComment          string
	GenerationExpression   string
}

func transColumn(colfs []columnstruct) (cols []column) {
	for _, colf := range colfs {
		var col column

		col.TableCatalog = tinyutils.TransValue(colf.TableCatalog)
		col.TableSchema = tinyutils.TransValue(colf.TableSchema)
		col.TableName = tinyutils.TransValue(colf.TableName)
		col.ColumnName = tinyutils.TransValue(colf.ColumnName)
		col.OrdinalPosition = tinyutils.TransValue(colf.OrdinalPosition)
		col.ColumnDefault = tinyutils.TransValue(colf.ColumnDefault)
		col.IsNullAble = tinyutils.TransValue(colf.IsNullAble)
		col.DataType = tinyutils.TransValue(colf.DataType)
		col.CharacterMaximumLength = tinyutils.TransValue(colf.CharacterMaximumLength)
		col.CharacterOrtetLength = tinyutils.TransValue(colf.CharacterOrtetLength)
		col.NumericPrecision = tinyutils.TransValue(colf.NumericPrecision)
		col.NumericScale = tinyutils.TransValue(colf.NumericScale)
		col.DataTimePrecision = tinyutils.TransValue(colf.DataTimePrecision)
		col.CharacterSetName = tinyutils.TransValue(colf.CharacterSetName)
		col.CollationName = tinyutils.TransValue(colf.CollationName)
		col.ColumnType = tinyutils.TransValue(colf.ColumnType)
		col.ColumnKey = tinyutils.TransValue(colf.ColumnKey)
		col.Extra = tinyutils.TransValue(colf.Extra)
		col.Privileges = tinyutils.TransValue(colf.Privileges)
		col.ColumnComment = tinyutils.TransValue(colf.ColumnComment)
		col.GenerationExpression = tinyutils.TransValue(colf.GenerationExpression)
		cols = append(cols, col)
	}
	return
}
