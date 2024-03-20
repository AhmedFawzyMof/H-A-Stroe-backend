package routes

import (
	"HAstore/database"
	"HAstore/middleware"
	"HAstore/models"
	"encoding/json"
	"net/http"
	"sync"
)

func NavData(res http.ResponseWriter, req *http.Request, params map[string]string) {
	res.WriteHeader(http.StatusOK)

	db := database.Connect()

	defer db.Close()

	var Category models.Category
	var Tag models.Tag

	wg := &sync.WaitGroup{}

	categoryChan := make(chan []byte, 1)
	tagChan := make(chan []byte, 1)

	wg.Add(2)
	go models.Category.GetAllCategories(Category, db, categoryChan, wg, false)
	go models.Tag.GetAllTags(Tag, db, tagChan, wg)
	wg.Wait()

	close(categoryChan)
	close(tagChan)

	var Categories []models.Category
	var Tags []models.Tag

	if err := json.Unmarshal(<-categoryChan, &Categories); err != nil {
		middleware.SendError(err, res)
		return
	}

	if err := json.Unmarshal(<-tagChan, &Tags); err != nil {
		middleware.SendError(err, res)
		return
	}

	Response := make(map[string]interface{}, 2)
	Response["Categories"] = Categories
	Response["Tags"] = Tags

	if err := json.NewEncoder(res).Encode(Response); err != nil {
		middleware.SendError(err, res)
		return
	}
}
