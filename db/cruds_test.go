package db

import (
	"math/rand"
	"testing"
	"time"

	"github.com/deep0ne/firstWBProject/models"
	"github.com/deep0ne/firstWBProject/utils"
	"github.com/stretchr/testify/require"
)

func createRandomOrder(t *testing.T) models.Order {
	orderInfo := models.OrderInfo{
		OrderUID:        utils.RandomString(20),
		TrackNumber:     utils.RandomString(15),
		Entry:           utils.RandomString(6),
		Delivery:        utils.RandomJSONText(5),
		Items:           utils.RandomItems(2),
		Locale:          utils.RandomString(3),
		InternalSign:    utils.RandomString(1),
		CustomerID:      utils.RandomString(5),
		DeliveryService: utils.RandomString(3),
		ShardKey:        utils.RandomString(5),
		SmID:            rand.Intn(15),
		DateCreated:     time.Now(),
		OofShard:        utils.RandomString(5),
	}
	paymentInfo := models.PaymentInfo{
		TransactionID: orderInfo.OrderUID,
		RequestID:     utils.RandomString(5),
		Currency:      utils.RandomString(3),
		Provider:      utils.RandomString(5),
		Amount:        rand.Intn(1000),
		PaymentDt:     1637907766,
		Bank:          utils.RandomString(5),
		DeliveryCost:  rand.Intn(50),
		GoodsTotal:    rand.Intn(100),
		CustomFee:     rand.Intn(5),
	}

	order := models.Order{OrderInfo: orderInfo, Payment: models.Payment{PaymentInfo: paymentInfo}}
	return order

}

func InsertAndGetRandomOrder(t *testing.T) {
	order := createRandomOrder(t)
	err := testDBWithCache.InsertIntoOrderInfo(order.OrderInfo)
	require.NoError(t, err)

	err = testDBWithCache.InsertIntoPaymentInfo(order.Payment)
	require.NoError(t, err)

	testOrder, err := testDBWithCache.SelectOrder(order.OrderUID)
	require.NoError(t, err)
	require.NotEmpty(t, testOrder)
	require.NotEmpty(t, testOrder.OrderUID)
	require.NotEmpty(t, testOrder.TransactionID)

	require.EqualValues(t, order, testOrder)
}
