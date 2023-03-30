package utils

import (
	"encoding/json"
	"math/rand"
	"strings"

	"github.com/jmoiron/sqlx/types"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyz"

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomJSONText(n int) types.JSONText {
	m := make(map[string]any)
	for i := 0; i < n; i++ {
		m[RandomString(n)] = RandomString(n)
	}
	jsonString, _ := json.Marshal(m)
	return jsonString
}

func RandomItems(items int) types.JSONText {
	m := make([]map[string]any, 0)
	for i := 0; i < items; i++ {
		lowerMap := make(map[string]any)
		item := RandomJSONText(5)
		json.Unmarshal(item, &lowerMap)
		m = append(m, lowerMap)
	}
	jsonString, _ := json.Marshal(m)
	return jsonString
}
