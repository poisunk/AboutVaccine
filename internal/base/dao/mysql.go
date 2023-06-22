package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"vax/internal/config"
	"xorm.io/xorm"
)

type DB struct {
	*xorm.Engine
}

func NewDB(engine *xorm.Engine) *DB {
	return &DB{engine}
}

func (db *DB) Close() error {
	return db.Engine.Close()
}

func NewEngine() (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("mysql", config.DSN)
	if err != nil {
		return nil, err
	}
	err = engine.Ping()
	if err != nil {
		return nil, err
	}
	return engine, nil
}
