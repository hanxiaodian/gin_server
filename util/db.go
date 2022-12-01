package util

import (
	"database/sql"
	"fmt"

	"gin_server/conf/setting"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (db *sql.DB) {
	conf := setting.Conf()
	dbConfig := conf.DataBase
	// 获取 MySQL 链接需要自己导入 _ "github.com/go-sql-driver/mysql"
	database, err := sql.Open("mysql", dbConfig.MYSQL_USERNAME+":"+dbConfig.MYSQL_PASSWORD+"@tcp("+dbConfig.MYSQL_WRITER_HOST+":"+dbConfig.MYSQL_WRITER_PORT+")/"+dbConfig.MYSQL_DATABASE)
	if err != nil {
		fmt.Println("open mysql failed,", err)
		defer database.Close()
		return
	}

	fmt.Printf("连接主数据库 port: %s;name: %s", dbConfig.MYSQL_WRITER_PORT, dbConfig.MYSQL_DATABASE)

	return database
	// 关闭数据库连接
	// defer database.Close() // 注意这行代码要写在上面err判断的下面
}
