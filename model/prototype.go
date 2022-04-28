package model

import (
	"github.com/jmoiron/sqlx"
)

type Engine struct {
	*sqlx.DB
}

func NewEngine(db *sqlx.DB) *Engine {
	return &Engine{
		DB: db,
	}
}
