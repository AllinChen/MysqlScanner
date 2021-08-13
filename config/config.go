package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/romberli/log"
	"gopkg.in/yaml.v3"
)

type Config struct {
	DB struct {
		Address     string `yaml:"address"`
		Port        int    `yaml:"port"`
		User        string `yaml:"user"`
		Password    string `yaml:"password"`
		DBName      string `yaml:"db_name"`
		MaxIdleConn int    `yaml:"max_idle_conn"`
		MaxOpenConn int    `yaml:"max_open_conn"`
	} `yaml:"db"`
}

var Conf *Config

const (
	configPath = "./config/config.yml"
)

var add = flag.String("add", "localhost", "数据库地址")
var port = flag.Int("port", 3306, "端口")
var user = flag.String("u", "", "用户名")
var pwd = flag.String("p", "", "用户密码")
var Fdb = flag.String("db", "", "指定数据库")
var Ftb = flag.String("tb", "", "指定数据表")

// initConfig 初始化读取配置文件，传递到全局变量Conf中
func InitConfig(path string) {
	if path == "" {
		path = configPath
	}

	conf, err := newConfig(path)
	if err != nil {
		log.Fatalf("init config failed: ", err)
	}
	conf.DB.Address = *add
	conf.DB.Port = *port
	conf.DB.User = *user
	conf.DB.Password = *pwd

	if *user == "" || *pwd == "" {
		fmt.Println("user or password cant be null")
		panic("error,user or password cant be null")
	}

	Conf = conf
}

// newConfig 根据文件地址读取配置内容
func newConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	err = yaml.NewDecoder(file).Decode(&config)

	return config, err
}
