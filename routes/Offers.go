package routes

import (
	"HAstore/database"
	"HAstore/middleware"
	"HAstore/models"
	"encoding/json"
	"net/http"
	"sync"
)

func Offers(res http.ResponseWriter, req *http.Request, params map[string]string) {
	res.WriteHeader(http.StatusOK)
	db := database.Connect()

	defer db.Close()

	var Offer models.Offers

	wg := &sync.WaitGroup{}

	offersChan := make(chan []byte, 1)

	wg.Add(1)
	go models.Offers.GetAllOffers(Offer, db, offersChan, wg)
	wg.Wait()

	close(offersChan)

	var Offers []models.Offers

	if err := json.Unmarshal(<-offersChan, &Offers); err != nil {
		middleware.SendError(err, res)
		return
	}

	Response := make(map[string]interface{})
	Response["Offers"] = Offers

	if err := json.NewEncoder(res).Encode(Response); err != nil {
		middleware.SendError(err, res)
		return
	}
}
