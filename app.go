package main

import (
	"database/sql"
	"log"

	"github.com/cbot918/youtube/auth/api"
)

const (
	port       = ":3005"
	staticPath = "ui/dist"
	dbtype     = "sqlite3"
	dbsource   = "auth.db"
)

func GetConfig() *api.Config {
	return &api.Config{
		PORT:       port,
		StaticPath: staticPath,
		DbType:     dbtype,
		DbSource:   dbsource,
	}
}

func GetDB(c *api.Config) *sql.DB {
	db, err := sql.Open(c.DbType, c.DbSource)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func GetDAO(db *sql.DB) *api.Dao {
	return &api.Dao{
		DB: db,
	}
}

func main() {

	cfg := GetConfig()
	dao := GetDAO(GetDB(cfg))

	app := api.New(cfg, dao)

	app.Run()
}
