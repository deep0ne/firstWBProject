package utils

import (
	"encoding/json"
	"time"

	"github.com/jmoiron/sqlx/types"
)

func UnmarshalDelivery(delivery types.JSONText) map[string]any {
	byteDelivery := types.JSONText([]byte(delivery))
	mapToPrint := make(map[string]any)
	json.Unmarshal(byteDelivery, &mapToPrint)
	return mapToPrint
}

func UnmarshalItems(items types.JSONText) []map[string]any {
	byteItems := types.JSONText([]byte(items))
	mapToPrint := make([]map[string]any, 0)
	json.Unmarshal(byteItems, &mapToPrint)
	return mapToPrint
}

func ParseUnix(unixTime int64) string {
	t := time.Unix(unixTime, 0)
	return t.Format("02-01-2006")
}
