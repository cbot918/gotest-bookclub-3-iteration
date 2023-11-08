package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type SQL interface {
	Connect() (*sql.DB, error)
}

type MySQL struct {
	driver string
	dsn    string
}

func (m MySQL) Connect() (*sql.DB, error) {
	db, err := sql.Open(m.driver, m.dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

type Postgres struct {
	driver string
	dsn    string
}

func (p Postgres) Connect() (*sql.DB, error) {
	db, err := sql.Open(p.driver, p.dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
