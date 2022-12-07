package db

import (
	"fmt"
	"gin_server/conf/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(dbConfig setting.Database) (db *gorm.DB) {
	var err error
	var database *gorm.DB

	sqlStr := dbConfig.MYSQL_USERNAME +
		":" + dbConfig.MYSQL_PASSWORD +
		"@tcp(" + dbConfig.MYSQL_WRITER_HOST +
		":" + dbConfig.MYSQL_WRITER_PORT + ")/" +
		dbConfig.MYSQL_DATABASE +
		"?charset=utf8mb4&parseTime=true&loc=Local"

	//配置项中预设了连接池 ConnPool
	database, err = gorm.Open(mysql.Open(sqlStr), &gorm.Config{})
	if err != nil {
		fmt.Println("connect mysql failed: ", err)
		return
	}

	fmt.Printf("connect mysql success port: %s;name: %s\n", dbConfig.MYSQL_WRITER_PORT, dbConfig.MYSQL_DATABASE)
	return database
}
