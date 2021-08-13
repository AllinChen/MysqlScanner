package markdown

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/AllinChen/MysqlScanner/tablescanner"
)

func BigTitle(content, databaseName string) string {

	return content + "# DATABASE - " + databaseName + " \n"
}

func SmallTitle(content, tableName string) string {
	return content + "## TABLE - " + tableName + " \n"
}

func DatabaseTable(content string, dbName string, tbs []tablescanner.Table) string {
	tbcontent := fmt.Sprintf("### 库 %s 中表信息", dbName) + "\n" +
		"| 序号 | 名称 | 来源 | 类型 | 存储引擎 | 创建时间 | 修改时间 | 备注 | \n" +
		"| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |\n"

	for nub, tb := range tbs {
		ct := fmt.Sprintf("| %d | %s | %s | %s | %s | %s | %s | %s |\n",
			nub+1,
			tb.TableName,
			tb.TableSchema,
			tb.TableType,
			tb.Engine,
			tb.CreateTime,
			tb.UpdateTime,
			tb.TableComment,
		)
		tbcontent += ct
	}

	if len(tbs) < 1 {
		ct := fmt.Sprintf("| %d | %s | %s | %s | %s | %s | %s | %s |\n",
			0,
			"",
			"",
			"",
			"",
			"",
			"",
			"",
		)
		tbcontent += ct
	}

	return content + tbcontent + "\n\n"
}

func TableColumn(content string, tbName string, cols []tablescanner.Column) string {
	colcontent := fmt.Sprintf(" 表 %s 中列信息", tbName) + "\n" +
		"| 序号 | 列名 | 表名 | 默认值 | 能否null | 数据类型 | 最大长度 | 列类型 | 其他 | 权限 | 备注 | \n" +
		"| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |\n"

	for _, col := range cols {
		ct := fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s |\n",
			col.OrdinalPosition,
			col.ColumnName,
			col.TableName,
			col.ColumnDefault,
			col.IsNullAble,
			col.DataType,
			col.CharacterMaximumLength,
			col.ColumnType,
			col.Extra,
			col.Privileges,
			col.ColumnComment,
		)
		colcontent += ct
	}
	return content + colcontent + "\n\n"
}

// 获取文件路径
func GetFilePath() (string, error) {
	pwd, err := os.Getwd()
	if err == nil {
		switch runtime.GOOS {
		case "darwin":
			return pwd + "/", nil
		case "windows":
			return pwd + "\\", nil
		case "linux":
			return pwd + "/", nil
		default:
			return "", nil
		}
	}

	return "", err
}

func SaveAsMdFile(path, content string) error {

	newFile, err := os.Create(path)
	defer newFile.Close()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, []byte(content), 0666)
	if err != nil {
		return err
	}
	return nil
}

func ReadFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	fmt.Println(string(data))
	return string(data), err
}
