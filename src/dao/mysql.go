package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

func InitMysql() {
	dsn := "root:zxc?12345@tcp(sh-cynosdbmysql-grp-o5q6mlei.sql.tencentcdb.com:22106)/vaccine?parseTime=True"
	var err error
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	DB.SingularTable(true)
	if err = DB.DB().Ping(); err != nil {
		panic(err)
	}
}
