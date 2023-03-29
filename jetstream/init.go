package jetstream

import (
	"log"

	"github.com/deep0ne/firstWBProject/models"
)

func JetStreamLaunch() []models.Order {
	log.Println("Starting JetStream...")
	js, err := JetStreamInit()
	checkErr(err)
	err = CreateStream(js)
	checkErr(err)
	err = CreateOrders(js)
	checkErr(err)
	orders, err := ConsumeOrders(js)
	checkErr(err)
	return orders
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
