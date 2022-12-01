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
	// 服务配置
	Project  project  `yaml:"PROJECT"`
	DataBase database `yaml:"DATABASE"`
	Redis    redis    `yaml:"REDIS"`
	Jwt      jwt      `yaml:"JWT"`
	SMS      sms      `yaml:"SMS_TEMPLATE"`
}

type project struct {
	// 项目配置
	APP_PORT      string `yaml:"APP_PORT"`
	APP_ENV       string `yaml:"APP_ENV"`
	PROJECT_NAME  string `yaml:"PROJECT_NAME"`
	WRITE_SWAGGER string `yaml:"WRITE_SWAGGER"`
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
	REDIS_PORT     string `yaml:"REDIS_PORT"`
	REDIS_PASSWORD string `yaml:"REDIS_PASSWORD"`
	REDIS_DB       int    `yaml:"REDIS_DB"`
}

type sms struct {
	TENCENT_CAPTCHA_SECRET_KEY string `yaml:"TENCENT_CAPTCHA_SECRET_KEY"`
	TENCENT_CAPTCHA_SECRET_ID  string `yaml:"TENCENT_CAPTCHA_SECRET_ID"`
	TENCENT_CAPTCHA_APP_ID     string `yaml:"TENCENT_CAPTCHA_APP_ID"`
	TENCENT_CAPTCHA_APP_SECRET string `yaml:"TENCENT_CAPTCHA_APP_SECRET"`
	MONTNETS_SMS_USER_ID       string `yaml:"MONTNETS_SMS_USER_ID"`
	MONTNETS_SMS_PWD           string `yaml:"MONTNETS_SMS_PWD"`
	MONTNETS_SMS_API_KEY       string `yaml:"MONTNETS_SMS_API_KEY"`
	MONTNETS_SMS_HOST          string `yaml:"MONTNETS_SMS_HOST"`
	MONTNETS_SMS_PORT          string `yaml:"MONTNETS_SMS_PORT"`
}

type jwt struct {
	// JWT 配置
	REDIS_HOST     string `yaml:"JWT_SECRET"`
	REDIS_PORT     string `yaml:"JWT_COOKIE_VERIFY"`
	REDIS_PASSWORD string `yaml:"COOKIE_SALT"`
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
		fmt.Println("ioutil.ReadFile failed: ", err)
	}

	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		// 获取错误处理
		fmt.Println("yaml.Unmarshal failed: ", err)
	}
}

// 获取配置  外部调用使用
func Conf() *Setting {
	return conf
}
