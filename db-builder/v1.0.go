package main

import (
	"github.com/jmoiron/sqlx"
)

const (
	DefCreateTableEmail string = `CREATE TABLE email (
		id BINARY(16) NOT NULL PRIMARY KEY,
		ctime TIMESTAMP NOT NULL DEFAULT current_timestamp,
		mtime TIMESTAMP NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
		enabled TINYINT(1) NOT NULL DEFAULT 1,
		addr VARCHAR(255) NOT NULL,
		password BINARY(16) NOT NULL,
		salt BINARY(16) NOT NULL,
		daily_count SMALLINT(4) NOT NULL DEFAULT 0,

		UNIQUE addr (addr),
		INDEX daily_count (daily_count)
	) ENGINE InnoDB COLLATE 'utf8mb4_unicode_ci' CHARACTER SET 'utf8mb4';`
)

func V1_0(db *sqlx.DB) error {
	db.MustExec(DefCreateTableEmail)
	return nil
}
