package db

import (
	"log"
	"os"
	"testing"

	"github.com/deep0ne/firstWBProject/utils"
	"github.com/jmoiron/sqlx"
)

var testDBWithCache Database
var testDB *sqlx.DB

func TestMain(m *testing.M) {
	cfg := utils.NewConfig()
	testDB, err := sqlx.Connect(cfg.DBDriver, cfg.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to the database:", err)
	}
	testDBWithCache = NewDBWithCache(testDB)

	os.Exit(m.Run())
}
