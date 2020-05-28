package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // 注册 mysql 驱动
)

var (
	// DB 导出的 DB
	DB *sql.DB
)

// ConnectionString 连接字符串
// user:password@tcp(127.0.0.1:1306)/testdb?charset=utf8mb4&parseTime=True&loc=Local
type ConnectionString struct {
	User     string
	Password string
	Host     string
	Port     int64
	Database string
	Charset  string
}

//resolveConnStr 返回 mysql 数据库连接字符串
func resolveConnStr(cs ConnectionString) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", cs.User, cs.Password, cs.Host, cs.Port, cs.Database, cs.Charset)
}

// InitMysql 配置 mysql 数据库
func InitMysql(cs ConnectionString, conns ...int) {
	connStr := resolveConnStr(cs)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}

	if len(conns) > 0 {
		db.SetMaxIdleConns(conns[0])
	}
	if len(conns) > 1 {
		db.SetMaxOpenConns(conns[1])
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	DB = db
}
