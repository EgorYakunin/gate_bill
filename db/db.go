package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Conn *sqlx.DB

func Connect() {
	connStr := "user=egor password=password123 dbname=gate_bill_db sslmode=disable"
	var err error
	Conn, err = sqlx.Connect("postgres", connStr)

	if err != nil {
		panic("db: can't open sql connection")
	}

	err = Conn.Ping()
	if err != nil {
		panic("db: can't ping a database")
	}
}

