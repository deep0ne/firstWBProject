package main

import (
	"log"

	"github.com/deep0ne/firstWBProject/api"
	"github.com/deep0ne/firstWBProject/db"
	"github.com/deep0ne/firstWBProject/jetstream"
	"github.com/jmoiron/sqlx"
)

const (
	dbDriver      = "pgx"
	postgreSQL    = "postgresql://root:wbpass@localhost:5432/wborders?sslmode=disable"
	serverAddress = "0.0.0.0:8081"
)

func main() {

	orders := jetstream.JetStreamLaunch()
	store, err := sqlx.Connect(dbDriver, postgreSQL)
	if err != nil {
		log.Fatal(err)
	}

	dbWithCache := db.NewDBWithCache(store)
	err = dbWithCache.RecoverCache()
	if err != nil {
		log.Println(err)
	}
	dbWithCache.FillDataIntoDB(orders)

	server, err := api.NewServer(dbWithCache)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
