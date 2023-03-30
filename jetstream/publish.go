package jetstream

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"github.com/deep0ne/firstWBProject/models"
	"github.com/nats-io/nats.go"
)

const (
	subjectName = "ORDERS.created"
)

func getOrders() ([]models.Order, error) {
	rawOrders, _ := ioutil.ReadFile("./model.json")
	var ordersObj []models.Order
	err := json.Unmarshal(rawOrders, &ordersObj)

	return ordersObj, err
}

func CreateOrders(js nats.JetStreamContext) error {
	// если не получилось санмаршалить в структуру, значит в файле лежало что-то не то
	// возвращаем ошибку
	orders, err := getOrders()
	if err != nil {
		return err
	}

	for _, oneOrder := range orders {

		// create random message intervals to slow down
		r := rand.Intn(1500)
		time.Sleep(time.Duration(r) * time.Millisecond)

		orderJSON, err := json.Marshal(oneOrder)
		if err != nil {
			log.Println(err)
			continue
		}

		_, err = js.Publish(subjectName, orderJSON)
		if err != nil {
			return err
		}
		log.Printf("Order with UID:%s has been published\n", oneOrder.OrderUID)
	}
	return nil
}
