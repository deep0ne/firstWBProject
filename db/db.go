package db

import (
	"github.com/deep0ne/firstWBProject/models"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	db    *sqlx.DB
	cache map[string]models.Order
}
