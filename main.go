package main

import (
	"database/sql"
	"log"

	"github.com/cbot918/youtube/auth/server"
)

const (
	port       = ":3005"
	staticPath = "ui/dist"
	dbtype     = "sqlite3"
	dbsource   = "auth.db"
)

func GetConfig() *server.Config {
	return &server.Config{
		PORT:       port,
		StaticPath: staticPath,
		DbType:     dbtype,
		DbSource:   dbsource,
	}
}

func GetDB(c *server.Config) *sql.DB {
	db, err := sql.Open(c.DbType, c.DbSource)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func GetDAO(db *sql.DB) *server.Dao {
	return &server.Dao{
		DB: db,
	}
}

func main() {

	cfg := GetConfig()
	dao := GetDAO(GetDB(cfg))

	app := server.New(cfg, dao)

	app.Run()
}
