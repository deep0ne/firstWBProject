package main

import (
	"log"

	"github.com/deep0ne/firstWBProject/api"
	"github.com/deep0ne/firstWBProject/db"
	"github.com/deep0ne/firstWBProject/jetstream"
	"github.com/deep0ne/firstWBProject/utils"
	"github.com/jmoiron/sqlx"
)

func main() {

	orders := jetstream.JetStreamLaunch()
	cfg := utils.NewConfig()
	store, err := sqlx.Connect(cfg.DBDriver, cfg.DBSource)
	if err != nil {
		log.Fatal(err)
	}

	defer store.Close()

	dbWithCache := db.NewDBWithCache(store)
	err = dbWithCache.RecoverCache()
	if err != nil {
		log.Println(err)
	}
	dbWithCache.FillDataIntoDB(orders)

	server, err := api.NewServer(dbWithCache)
	err = server.Start(cfg.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
