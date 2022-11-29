package setting

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// 相应设置配置
type Setting struct {
	RunMode string `yaml:"runMode"`
	// 服务配置
	Project project `yaml:"project"`

	DataBase database `yaml:"database"`

	Redis redis `yaml:"redis"`
}

type project struct {
	// 项目配置
	APP_PORT     string `yaml:"APP_PORT"`
	APP_ENV      string `yaml:"APP_ENV"`
	PROJECT_NAME string `yaml:"PROJECT_NAME"`
}

type database struct {
	// 数据库配置
	MYSQL_USERNAME    string `yaml:"MYSQL_USERNAME"`
	MYSQL_PASSWORD    string `yaml:"MYSQL_PASSWORD"`
	MYSQL_WRITER_PORT string `yaml:"MYSQL_WRITER_PORT"`
	MYSQL_WRITER_HOST string `yaml:"MYSQL_WRITER_HOST"`
	MYSQL_READER_PORT string `yaml:"MYSQL_READER_PORT"`
	MYSQL_READER_HOST string `yaml:"MYSQL_READER_HOST"`
	MYSQL_DATABASE    string `yaml:"MYSQL_DATABASE"`
}

type redis struct {
	// redis配置
	REDIS_HOST     string `yaml:"REDIS_HOST"`
	REDIS_PORT     int    `yaml:"REDIS_PORT"`
	REDIS_PASSWORD string `yaml:"REDIS_PASSWORD"`
	REDIS_DB       int    `yaml:"REDIS_DB"`
}

var conf = &Setting{}

// 初始化方法
func InitSetting() {
	serverPath, err := os.Getwd() // 获得根目录文件路径
	if err != nil {
		// 错误处理
		fmt.Println("os.Getwd    failed")
	}

	yamlFile, err := ioutil.ReadFile(filepath.Join(serverPath, "./config.local.yaml"))
	if err != nil {
		// 错误处理
		fmt.Println("ioutil.ReadFile    failed")
	}

	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		// 获取错误处理
		fmt.Println("yaml.Unmarshal    failed")
	}
}

// 获取配置  外部调用使用
func Conf() *Setting {
	return conf
}
