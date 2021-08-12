package config

import (
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

// initConfig 初始化读取配置文件，传递到全局变量Conf中
func InitConfig(path string) {
	if path == "" {
		path = configPath
	}

	conf, err := newConfig(path)
	if err != nil {
		log.Fatalf("init config failed: ", err)
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
