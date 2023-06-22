package dao

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"vax/internal/config"
	"vax/internal/entity"
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

var tables = []interface{}{
	&entity.AdverseEvent{},
	&entity.AdverseSymptom{},
	&entity.AdverseVaccine{},
	&entity.OAETerm{},
	&entity.User{},
	&entity.Vaccine{},
	&entity.VaccineType{},
	&entity.Vaers{},
	&entity.VaersResult{},
	&entity.VaersSymptom{},
	&entity.VaersSymptomTerm{},
	&entity.VaersVax{},
	&entity.VaersVaxTerm{},
}

func NewEngine(data *config.Database) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine(data.Driver, data.Connection)
	if err != nil {
		return nil, err
	}
	err = engine.Ping()
	if err != nil {
		return nil, err
	}
	err = engine.Sync(tables)
	if err != nil {
		return nil, err
	}
	return engine, nil
}
