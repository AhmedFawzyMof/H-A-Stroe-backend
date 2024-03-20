package routes

import (
	"HAstore/database"
	"HAstore/middleware"
	"HAstore/models"
	"encoding/json"
	"net/http"
	"sync"
)

func AllCategories(res http.ResponseWriter, req *http.Request, params map[string]string) {
	res.WriteHeader(http.StatusOK)

	db := database.Connect()

	defer db.Close()

	var Category models.Category

	wg := &sync.WaitGroup{}

	categoryChan := make(chan []byte, 1)

	wg.Add(1)
	go models.Category.GetAllCategories(Category, db, categoryChan, wg, true)
	wg.Wait()

	close(categoryChan)

	var Categories []models.Category

	if err := json.Unmarshal(<-categoryChan, &Categories); err != nil {
		middleware.SendError(err, res)
		return
	}

	Response := make(map[string]interface{})
	Response["Categories"] = Categories

	if err := json.NewEncoder(res).Encode(Response); err != nil {
		middleware.SendError(err, res)
		return
	}
}
