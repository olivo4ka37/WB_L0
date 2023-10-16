package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
	"github.com/olivo4ka37/WB_L0/internal/models"
	"log"
	"math/rand"
	"time"
)

func RandString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
func RandNumberToString(n int) string {
	var letterRunes = []rune("0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	nc, err := stan.Connect("test-cluster", "publisher", stan.NatsURL("0.0.0.0:4222"))
	if err != nil {
		log.Fatal("cant connect to nats-streaming", err)
		return
	}
	defer nc.Close()

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		id := uuid.New()
		order := models.Order{
			OrderUID:    id.String(),
			TrackNumber: RandNumberToString(10),
			Entry:       RandString(4),
			Delivery: models.Delivery{
				Name:    RandString(20),
				Phone:   "+" + RandNumberToString(11),
				Zip:     RandString(7),
				City:    RandString(30),
				Address: RandString(30),
				Region:  RandString(30),
				Email:   RandString(10) + "gmail.com",
			},
			Payment: models.Payment{
				Transaction:  id.String(),
				RequestID:    RandNumberToString(10),
				Currency:     RandNumberToString(3),
				Provider:     RandNumberToString(3),
				Amount:       123,
				PaymentDt:    123,
				Bank:         RandString(10),
				DeliveryCost: 123,
				GoodsTotal:   123,
				CustomFee:    123,
			},
			Items: []models.Item{
				{
					ChrtID:      12345,
					TrackNumber: RandString(20),
					Price:       12345,
					Rid:         RandString(20),
					Name:        RandString(20),
					Sale:        1234567,
					Size:        RandString(20),
					TotalPrice:  1,
					NmID:        21454,
					Brand:       RandString(20),
					Status:      1234,
				},
			},
			Locale:            RandString(20),
			InternalSignature: RandString(20),
			CustomerID:        RandString(20),
			DeliveryService:   RandString(20),
			ShardKey:          RandString(20),
			SmID:              12323232,
			DateCreated:       time.Now(),
			OofShard:          RandString(20),
		}

		bytes, err := json.Marshal(order)
		if err != nil {
			log.Println(err)
			return
		}

		err = nc.Publish("order", bytes)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
