package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/yogaprasetya22/simple-bank/api"
	db "github.com/yogaprasetya22/simple-bank/db/sqlc"
	"github.com/yogaprasetya22/simple-bank/util"
)


func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
