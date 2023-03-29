package models

import (
	"time"

	"github.com/jmoiron/sqlx/types"
)

type OrderInfo struct {
	OrderUID        string         `json:"order_uid" db:"order_uid"`
	TrackNumber     string         `json:"track_number" db:"track_number"`
	Entry           string         `json:"entry" db:"entry"`
	Delivery        types.JSONText `json:"delivery" db:"delivery"`
	Items           types.JSONText `json:"items" db:"items"`
	Locale          string         `json:"locale" db:"locale"`
	InternalSign    string         `json:"internal_signature" db:"internal_signature"`
	CustomerID      string         `json:"customer_id" db:"customer_id"`
	DeliveryService string         `json:"delivery_service" db:"delivery_service"`
	ShardKey        string         `json:"shardkey" db:"shardkey"`
	SmID            int            `json:"sm_id" db:"sm_id"`
	DateCreated     time.Time      `json:"date_created" db:"date_created"`
	OofShard        string         `json:"oof_shard" db:"oof_shard"`
}

type PaymentInfo struct {
	TransactionID string `json:"transaction" db:"transaction"`
	RequestID     string `json:"request_id" db:"request_id"`
	Currency      string `json:"currency" db:"currency"`
	Provider      string `json:"provider" db:"provider"`
	Amount        int    `json:"amount" db:"amount"`
	PaymentDt     int64  `json:"payment_dt" db:"payment_dt"`
	Bank          string `json:"bank" db:"bank"`
	DeliveryCost  int    `json:"delivery_cost" db:"delivery_cost"`
	GoodsTotal    int    `json:"goods_total" db:"goods_total"`
	CustomFee     int    `json:"custom_fee" db:"custom_fee"`
}

type Payment struct {
	PaymentInfo `json:"payment"`
}

type Order struct {
	OrderInfo
	Payment
}
