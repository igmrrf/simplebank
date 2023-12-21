package main

import (
	"database/sql"
	"log"

	"github.com/igmrrf/simplebank/api"
	db "github.com/igmrrf/simplebank/db/sqlc"
	"github.com/igmrrf/simplebank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.SERVERADDRESS)
	if err != nil {
		log.Fatal("Cannot run server: ", err)
	}
}
