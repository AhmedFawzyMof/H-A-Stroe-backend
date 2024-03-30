package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"sync"
)

type Offers struct {
	Id    	int    `json:"id"`
	Img 	string `json:"image"`
	Product string `json:"product"`
}

func (o Offers) GetAllOffers(db *sql.DB, offersChan chan []byte, wg *sync.WaitGroup) {
	defer wg.Done()

	var TOffers []Offers

	offers, err := db.Query("SELECT * FROM Offers")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer offers.Close()

	for offers.Next() {
		var offer Offers

		if offers.Scan(&offer.Id, &offer.Img,	&offer.Product); err != nil {
			fmt.Println(err.Error())
		}


		TOffers = append(TOffers, offer)
	}

	offersBytes, err := json.Marshal(TOffers)

	if err != nil {
		fmt.Println(err.Error())
	}

	offersChan <- offersBytes
}
