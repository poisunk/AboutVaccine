package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type DB struct {
	*xorm.Engine
}

const (
	DSN = "root:zxc?12345@tcp(sh-cynosdbmysql-grp-o5q6mlei.sql.tencentcdb.com:22106)/vaccine?parseTime=True"
)

func NewDB(engine *xorm.Engine) *DB {
	return &DB{engine}
}

func (db *DB) Close() error {
	return db.Engine.Close()
}

func NewEngine(dsn string) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = engine.Ping()
	if err != nil {
		return nil, err
	}
	return engine, nil
}
