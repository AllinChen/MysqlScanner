package tablescanner

// MysqlDefaultDatabase mysql中的默认表
var MysqlDefaultDatabase []string = []string{
	"information_schema",
	"mysql",
	"performance_schema",
	"sys",
}

// database
type Database struct {
	Database string `bdb:"Database"`
}

// GetDatabases 得到数据库的列表
func GetDatabases(defaultDatabase []string) ([]Database, error) {
	sql := "show databases;"

	db, err := DBTool.GetDBConnt()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var dbs []Database
	for rows.Next() {
		var db Database
		if err = rows.Scan(&db.Database); nil != err {
			return nil, err
		}
		_, found := find(defaultDatabase, db.Database)
		if !found {
			dbs = append(dbs, db)
		}

	}
	return dbs, nil

}

// find 封装一个find来对比列表中是否已存在的值
func find(something []string, wanted string) (int, bool) {
	for nb, item := range something {
		if wanted == item {
			return nb, true
		}
	}
	return -1, false
}
