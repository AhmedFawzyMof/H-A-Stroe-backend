package routes

import (
	"HAstore/database"
	"HAstore/middleware"
	"HAstore/models"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
)

func CheckOut(res http.ResponseWriter, req *http.Request, params map[string]string) {
	res.WriteHeader(http.StatusOK)

	db := database.Connect()

	defer db.Close()

	var Order models.Order

	if err := json.NewDecoder(req.Body).Decode(&Order); err != nil {
		middleware.SendError(err, res)
		return
	}

	orderId := uuid.New()

	Order.Id = fmt.Sprintf("%v", orderId)

	if Order.Method == "paypal" {
		Order.IsPaid = true
	}
	Order.CreatedAt = time.Now()

	wg := &sync.WaitGroup{}

	orderChan := make(chan []byte, 1)

	wg.Add(1)

	go models.Order.Create(Order, db, orderChan, wg)

	wg.Wait()

	close(orderChan)

	Response := make(map[string]interface{})

	if err := json.Unmarshal(<-orderChan, &Response); err != nil {
		middleware.SendError(err, res)
	}

	if err := json.NewEncoder(res).Encode(Response); err != nil {
		middleware.SendError(err, res)
	}
}

func AuthCheckOut(res http.ResponseWriter, req *http.Request, params map[string]string) {
	res.WriteHeader(http.StatusOK)

	db := database.Connect()

	defer db.Close()

	var Order models.Order

	if err := json.NewDecoder(req.Body).Decode(&Order); err != nil {
		middleware.SendError(err, res)
		return
	}

	Order.User = params["authEmail"]

	orderId := uuid.New()

	Order.Id = fmt.Sprintf("%v", orderId)

	if Order.Method == "paypal" {
		Order.IsPaid = true
	}
	Order.CreatedAt = time.Now()

	wg := &sync.WaitGroup{}

	orderChan := make(chan []byte, 1)

	wg.Add(1)

	go models.Order.Create(Order, db, orderChan, wg)

	wg.Wait()

	close(orderChan)

	Response := make(map[string]interface{})

	if err := json.Unmarshal(<-orderChan, &Response); err != nil {
		middleware.SendError(err, res)
	}

	if err := json.NewEncoder(res).Encode(Response); err != nil {
		middleware.SendError(err, res)
	}
}

func OrderHistroy(res http.ResponseWriter, req *http.Request, params map[string]string) {
	res.WriteHeader(http.StatusOK)

	db := database.Connect()

	defer db.Close()

	var Order models.Order

	Order.User = params["authEmail"]

	var TheOrders []models.Order = models.Order.GetHistory(Order, db)

	Orders := models.OrdersS(TheOrders)

	Orders.GetOrderProducts(db)

	var Response map[string]interface{} = map[string]interface{}{
		"Orders": Orders,
	}

	if err := json.NewEncoder(res).Encode(Response); err != nil {
		middleware.SendError(err, res)
	}
}
