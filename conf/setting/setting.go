package setting

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"

	"gopkg.in/yaml.v2"

	"gin_server/util"
)

type Env struct {
	APP_PORT                   string `yaml:"APP_PORT"`
	APP_ENV                    string `yaml:"APP_ENV"`
	PROJECT_NAME               string `yaml:"PROJECT_NAME"`
	WRITE_SWAGGER              string `yaml:"WRITE_SWAGGER"`
	MYSQL_USERNAME             string `yaml:"MYSQL_USERNAME"`
	MYSQL_PASSWORD             string `yaml:"MYSQL_PASSWORD"`
	MYSQL_WRITER_PORT          string `yaml:"MYSQL_WRITER_PORT"`
	MYSQL_WRITER_HOST          string `yaml:"MYSQL_WRITER_HOST"`
	MYSQL_READER_PORT          string `yaml:"MYSQL_READER_PORT"`
	MYSQL_READER_HOST          string `yaml:"MYSQL_READER_HOST"`
	MYSQL_DATABASE             string `yaml:"MYSQL_DATABASE"`
	REDIS_HOST                 string `yaml:"REDIS_HOST"`
	REDIS_PORT                 int32  `yaml:"REDIS_PORT"`
	REDIS_PASSWORD             string `yaml:"REDIS_PASSWORD"`
	REDIS_DB                   int    `yaml:"REDIS_DB"`
	TENCENT_CAPTCHA_SECRET_KEY string `yaml:"TENCENT_CAPTCHA_SECRET_KEY"`
	TENCENT_CAPTCHA_SECRET_ID  string `yaml:"TENCENT_CAPTCHA_SECRET_ID"`
	TENCENT_CAPTCHA_APP_ID     string `yaml:"TENCENT_CAPTCHA_APP_ID"`
	TENCENT_CAPTCHA_APP_SECRET string `yaml:"TENCENT_CAPTCHA_APP_SECRET"`
	MONTNETS_SMS_USER_ID       string `yaml:"MONTNETS_SMS_USER_ID"`
	MONTNETS_SMS_PWD           string `yaml:"MONTNETS_SMS_PWD"`
	MONTNETS_SMS_API_KEY       string `yaml:"MONTNETS_SMS_API_KEY"`
	MONTNETS_SMS_HOST          string `yaml:"MONTNETS_SMS_HOST"`
	MONTNETS_SMS_PORT          string `yaml:"MONTNETS_SMS_PORT"`
	JWT_SECRET                 string `yaml:"JWT_SECRET"`
	JWT_COOKIE_VERIFY          string `yaml:"JWT_COOKIE_VERIFY"`
	COOKIE_SALT                string `yaml:"COOKIE_SALT"`
}

// 相应设置配置
type Setting struct {
	// 服务配置
	Project  Project  `yaml:"PROJECT"`
	DataBase Database `yaml:"DATABASE"`
	Redis    Redis    `yaml:"REDIS"`
	Jwt      Jwt      `yaml:"JWT"`
	SMS      SMS      `yaml:"SMS_TEMPLATE"`
}

type Project struct {
	// 项目配置
	APP_PORT      string `yaml:"APP_PORT"`
	APP_ENV       string `yaml:"APP_ENV"`
	PROJECT_NAME  string `yaml:"PROJECT_NAME"`
	WRITE_SWAGGER string `yaml:"WRITE_SWAGGER"`
}

type Database struct {
	// 数据库配置
	MYSQL_USERNAME    string `yaml:"MYSQL_USERNAME"`
	MYSQL_PASSWORD    string `yaml:"MYSQL_PASSWORD"`
	MYSQL_WRITER_PORT string `yaml:"MYSQL_WRITER_PORT"`
	MYSQL_WRITER_HOST string `yaml:"MYSQL_WRITER_HOST"`
	MYSQL_READER_PORT string `yaml:"MYSQL_READER_PORT"`
	MYSQL_READER_HOST string `yaml:"MYSQL_READER_HOST"`
	MYSQL_DATABASE    string `yaml:"MYSQL_DATABASE"`
}

type Redis struct {
	// redis 配置
	REDIS_HOST     string `yaml:"REDIS_HOST"`
	REDIS_PORT     int32  `yaml:"REDIS_PORT"`
	REDIS_PASSWORD string `yaml:"REDIS_PASSWORD"`
	REDIS_DB       int    `yaml:"REDIS_DB"`
}

type SMS struct {
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

type Jwt struct {
	// JWT 配置
	JWT_SECRET        string `yaml:"JWT_SECRET"`
	JWT_COOKIE_VERIFY string `yaml:"JWT_COOKIE_VERIFY"`
	COOKIE_SALT       string `yaml:"COOKIE_SALT"`
}

var result = &Env{}

// 初始化方法
func InitSetting() {
	conf := &Env{}
	localConf := &Env{}
	testConf := &Env{}
	serverPath, err := os.Getwd() // 获得根目录文件路径
	if err != nil {
		fmt.Println("settings os.Getwd failed")
	}

	yamlFile, err := ioutil.ReadFile(filepath.Join(serverPath, "./config.yaml"))
	if err != nil {
		fmt.Println("ioutil.ReadFile failed: ", err)
	}

	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		fmt.Println("yaml.Unmarshal failed: ", err)
	}

	if util.Exists(filepath.Join(serverPath, "./config.local.yaml")) {
		yamlFile, err := ioutil.ReadFile(filepath.Join(serverPath, "./config.local.yaml"))
		if err != nil {
			// 错误处理
			fmt.Println("ioutil.ReadFile failed: ", err)
		}
		yaml.Unmarshal(yamlFile, localConf)
	}

	if util.Exists(filepath.Join(serverPath, "./config.test.yaml")) {
		yamlFile, err := ioutil.ReadFile(filepath.Join(serverPath, "./config.test.yaml"))
		if err != nil {
			// 错误处理
			fmt.Println("ioutil.ReadFile failed: ", err)
		}
		yaml.Unmarshal(yamlFile, testConf)
	}

	fields := reflect.ValueOf(*result)
	localValues := reflect.ValueOf(*localConf)
	testValues := reflect.ValueOf(*testConf)
	prodValues := reflect.ValueOf(*conf)
	values := reflect.ValueOf(result).Elem()

	// 遍历结构体所有成员
	for i := 0; i < fields.NumField(); i++ {
		localValue := localValues.Field(i)
		testValue := testValues.Field(i)
		value := prodValues.Field(i)
		// 优先读取 .local.yaml 的配置信息
		if !reflect.DeepEqual(localValue.Interface(), reflect.Zero(localValue.Type()).Interface()) {
			values.Field(i).Set(localValue)
			// fmt.Println("localValue:   ", localValue)
			continue
		}

		// 其次读取 .test.yaml 的配置信息
		if !reflect.DeepEqual(testValue.Interface(), reflect.Zero(testValue.Type()).Interface()) {
			values.Field(i).Set(testValue)
			// fmt.Println("testValue:   ", testValue)
			continue
		}
		values.Field(i).Set(value)
	}
}

// 获取配置  外部调用使用
func Conf() (config *Setting, env *Env) {
	conf := result
	setting := &Setting{
		Project: Project{
			APP_ENV:       conf.APP_ENV,
			APP_PORT:      conf.APP_PORT,
			PROJECT_NAME:  conf.PROJECT_NAME,
			WRITE_SWAGGER: conf.WRITE_SWAGGER,
		},
		DataBase: Database{
			MYSQL_USERNAME:    conf.MYSQL_USERNAME,
			MYSQL_PASSWORD:    conf.MYSQL_PASSWORD,
			MYSQL_WRITER_PORT: conf.MYSQL_WRITER_PORT,
			MYSQL_WRITER_HOST: conf.MYSQL_WRITER_HOST,
			MYSQL_READER_PORT: conf.MYSQL_READER_PORT,
			MYSQL_READER_HOST: conf.MYSQL_READER_HOST,
			MYSQL_DATABASE:    conf.MYSQL_DATABASE,
		},
		Redis: Redis{
			REDIS_DB:       conf.REDIS_DB,
			REDIS_HOST:     conf.REDIS_HOST,
			REDIS_PORT:     conf.REDIS_PORT,
			REDIS_PASSWORD: conf.REDIS_PASSWORD,
		},
		Jwt: Jwt{
			JWT_SECRET:        conf.JWT_SECRET,
			COOKIE_SALT:       conf.COOKIE_SALT,
			JWT_COOKIE_VERIFY: conf.JWT_COOKIE_VERIFY,
		},
		SMS: SMS{},
	}
	return setting, conf
}
