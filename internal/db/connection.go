package db

import (
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var schema string = "CREATE TABLE `auto` (" +
	"`id` int(10) unsigned NOT NULL AUTO_INCREMENT, " +
	"`brand` varchar(32) NOT NULL, " +
	"`model` varchar(255) NOT NULL, " +
	"`engine_volume` decimal(5,2) unsigned NOT NULL, " +
	"PRIMARY KEY (`id`), " +
	"UNIQUE KEY `id_UNIQUE` (`id`) " +
	") ENGINE=InnoDB DEFAULT CHARSET=latin1"

func Connect() *sqlx.DB {
	dbURL := os.Getenv("database_url")
	if dbURL == "" {
		return nil
	}
	conn, err := sqlx.Connect("mysql", dbURL)
	if err != nil {
		panic(err)
	}
	_, err = conn.Exec(schema)
	if err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			panic(err)
		}
	}
	return conn
}
