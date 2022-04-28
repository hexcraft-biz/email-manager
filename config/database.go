package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

//================================================================
//
//================================================================
type DBI struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	Params   string
	MaxOpen  int
	MaxIdle  int
	LifeTime int
	IdleTime int
}

func (dbi *DBI) Open() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", dbi.protocol())
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(dbi.MaxOpen)
	db.SetMaxIdleConns(dbi.MaxIdle)
	db.SetConnMaxLifetime(time.Duration(dbi.LifeTime) * time.Second)
	db.SetConnMaxIdleTime(time.Duration(dbi.IdleTime) * time.Second)

	return db, nil
}

func (dbi *DBI) protocol() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		dbi.User,
		dbi.Password,
		dbi.Host,
		dbi.Port,
		dbi.Name,
		dbi.Params,
	)
}
