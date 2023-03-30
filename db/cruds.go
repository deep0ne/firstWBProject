package db

import (
	"errors"
	"log"

	"github.com/deep0ne/firstWBProject/models"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewDBWithCache(db *sqlx.DB) Database {
	return Database{
		db:    db,
		cache: make(map[string]models.Order),
	}
}

func (db *Database) FillDataIntoDB(orders []models.Order) error {

	for _, order := range orders {
		db.InsertIntoCache(order)

		err := db.InsertIntoOrderInfo(order.OrderInfo)
		if err != nil {
			return err
		}

		err = db.InsertIntoPaymentInfo(order.Payment)
		if err != nil {
			return err
		}

	}

	return nil
}

func (db *Database) InsertIntoCache(order models.Order) {
	if _, ok := db.cache[order.OrderUID]; ok {
		return
	}
	db.cache[order.OrderUID] = order
	log.Println("Successfully cached data with orderUID: ", order.OrderUID)
}

func (db *Database) RecoverCache() error {
	var orders []models.Order
	err := db.db.Select(&orders, "SELECT * FROM order_info o FULL OUTER JOIN payment p ON o.order_uid = p.transaction")
	if len(orders) == 0 || err != nil {
		return errors.New("Database is empty!")
	}

	log.Println("Recovering cache from the database...")
	log.Println("Number of rows in cache:", len(db.cache))

	for _, order := range orders {
		db.cache[order.OrderUID] = order
	}
	log.Println("Number of rows in cache AFTER RECOVER:", len(db.cache))
	return nil
}

func (db *Database) InsertIntoOrderInfo(info models.OrderInfo) error {
	query := `INSERT INTO order_info(
		order_uid, track_number,
		entry, delivery, items,
		locale, internal_signature,
		customer_id, delivery_service,
		shardkey, sm_id,
		date_created, oof_shard) 
	VALUES(:order_uid, :track_number, :entry, :delivery, :items, :locale, :internal_signature, :customer_id, :delivery_service, :shardkey, :sm_id, :date_created, :oof_shard)`

	_, err := db.db.NamedExec(query, info)
	if err != nil {
		return err
	}
	log.Printf("Order with UID:%s is successfully inserted into PostgreSQL\n", info.OrderUID)
	return nil
}

func (db *Database) InsertIntoPaymentInfo(payment models.Payment) error {
	query := `INSERT INTO payment(
		transaction, request_id, currency,
		provider, amount, payment_dt, bank,
		delivery_cost, goods_total, custom_fee)
		VALUES(:transaction, :request_id, :currency, :provider, :amount, :payment_dt, :bank, :delivery_cost, :goods_total, :custom_fee)`

	_, err := db.db.NamedExec(query, payment.PaymentInfo)
	if err != nil {
		return err
	}
	log.Printf("Payment to order with UID:%s is successfully inserted into PostgreSQL\n", payment.PaymentInfo.TransactionID)
	return nil
}

func (db *Database) SelectOrder(orderUID string) (models.Order, error) {
	if val, ok := db.cache[orderUID]; ok {
		log.Println("Getting data from cache...")
		return val, nil
	}

	rows, err := db.db.Query("SELECT * FROM order_info o FULL OUTER JOIN payment p ON o.order_uid = p.transaction")
	if err != nil {
		return models.Order{}, err
	}

	var order models.Order
	for rows.Next() {
		err = rows.Scan(&order)
	}

	if err != nil {
		return models.Order{}, err
	}
	// Заполняем кеш
	db.cache[orderUID] = order
	return order, nil
}
