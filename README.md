# MysqlScanner
MysqlScanner
用于将数据库内的表结构信息输出为markdown文件

### 已支持数据库
    · MySQL
    · TiDB
## How to use

使用命令行进行使用
参数
    - add   数据库地址，默认localhost

    - port  端口号，默认3306

    - u     用户名

    - p     密码

    - db    指定数据库，默认为用户权限内的所有非系统库

    - tb    指定表，在未指定数据库的情况下，会输出用户权限所能接触到的所有同名表

    eg: go run main.go -add 127.0.0.1 -port 3306 -u root -p 1234 -db test -tb test

    成功后会在根目录生成对应markdown文件