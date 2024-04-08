package routes

import (
	"HAstore/database"
	"HAstore/middleware"
	"HAstore/models"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func GetProduct(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	
	id, err := strconv.Atoi(req.PathValue("id"))

	if err != nil {
		er := errors.New("invalid id")
		middleware.SendError(er, res)
		return
	}

	db := database.Connect()
	defer db.Close()

	product := models.Product{}
	Product, err := product.GetProductById(db, id)

	if err != nil {
		middleware.SendError(err, res)
		return
	}

	if err := json.NewEncoder(res).Encode(Product); err != nil {
		middleware.SendError(err, res)
		return
	}
}