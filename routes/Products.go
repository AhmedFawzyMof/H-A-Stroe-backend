package routes

import (
	"HAstore/database"
	"HAstore/middleware"
	"HAstore/models"
	"encoding/json"
	"io"
	"net/http"
	"sync"
)

func AllProducts(res http.ResponseWriter, req *http.Request, params map[string]string) {
	res.WriteHeader(http.StatusOK)

	db := database.Connect()

	defer db.Close()

	var Product models.Product

	wg := &sync.WaitGroup{}

	productChan := make(chan []byte, 1)

	wg.Add(1)
	go models.Product.GetAllProduct(Product, db, productChan, wg)
	wg.Wait()

	close(productChan)

	var Products []models.Product

	if err := json.Unmarshal(<-productChan, &Products); err != nil {
		middleware.SendError(err, res)
		return
	}

	Response := make(map[string]interface{}, 1)
	Response["Products"] = Products

	if err := json.NewEncoder(res).Encode(Response); err != nil {
		middleware.SendError(err, res)
		return
	}
}

func Filter(res http.ResponseWriter, req *http.Request, params map[string]string) {
	res.WriteHeader(http.StatusOK)
	BodyData := req.Body
	Data, err := io.ReadAll(BodyData)

	if err != nil {
		middleware.SendError(err, res)
		return
	}

	var filters models.FilterData

	if err := json.Unmarshal(Data, &filters); err != nil {
		middleware.SendError(err, res)
		return
	}

	db := database.Connect()

	defer db.Close()

	var Product models.Product

	wg := &sync.WaitGroup{}

	productChan := make(chan []byte, 1)

	wg.Add(1)
	go models.Product.FilteredProducts(Product, db, filters, productChan, wg)
	wg.Wait()

	close(productChan)

	var Products []models.Product

	if err := json.Unmarshal(<-productChan, &Products); err != nil {
		middleware.SendError(err, res)
		return
	}

	Response := make(map[string]interface{}, 1)
	Response["Products"] = Products

	if err := json.NewEncoder(res).Encode(Response); err != nil {
		middleware.SendError(err, res)
		return
	}

}

func ProductBySlug(res http.ResponseWriter, req *http.Request, params map[string]string) {
	res.WriteHeader(http.StatusOK)
	db := database.Connect()

	defer db.Close()

	var Product models.Product
	Product.Slug = params["slug"]

	wg := &sync.WaitGroup{}

	productChan := make(chan []byte, 1)

	wg.Add(1)
	go models.Product.ProductBySlug(Product, db, productChan, wg)
	wg.Wait()

	close(productChan)

	var Products []models.Product

	if err := json.Unmarshal(<-productChan, &Products); err != nil {
		middleware.SendError(err, res)
		return
	}

	Response := make(map[string]interface{}, 1)
	Response["Products"] = Products

	if err := json.NewEncoder(res).Encode(Response); err != nil {
		middleware.SendError(err, res)
		return
	}
}

func ProductsByCategory(res http.ResponseWriter, req *http.Request, params map[string]string) {
	res.WriteHeader(http.StatusOK)

	db := database.Connect()

	defer db.Close()

	var Product models.Product
	Product.Category = params["name"]

	wg := &sync.WaitGroup{}

	productChan := make(chan []byte, 1)

	wg.Add(1)
	go models.Product.ProductsByCategorys(Product, db, productChan, wg)
	wg.Wait()

	close(productChan)

	var Products []models.Product

	if err := json.Unmarshal(<-productChan, &Products); err != nil {
		middleware.SendError(err, res)
		return
	}

	Response := make(map[string]interface{}, 1)
	Response["Products"] = Products

	if err := json.NewEncoder(res).Encode(Response); err != nil {
		middleware.SendError(err, res)
		return
	}
}

func ProductsByTag(res http.ResponseWriter, req *http.Request, params map[string]string) {
	res.WriteHeader(http.StatusOK)

	db := database.Connect()

	defer db.Close()

	var Product models.Product
	Product.Tag = params["name"]

	wg := &sync.WaitGroup{}

	productChan := make(chan []byte, 1)

	wg.Add(1)
	go models.Product.ProductsByTag(Product, db, productChan, wg)
	wg.Wait()

	close(productChan)

	var Products []models.Product

	if err := json.Unmarshal(<-productChan, &Products); err != nil {
		middleware.SendError(err, res)
		return
	}

	Response := make(map[string]interface{}, 1)
	Response["Products"] = Products

	if err := json.NewEncoder(res).Encode(Response); err != nil {
		middleware.SendError(err, res)
		return
	}
}

func ProductsBySearch(res http.ResponseWriter, req *http.Request, params map[string]string) {
	res.WriteHeader(http.StatusOK)

	db := database.Connect()

	defer db.Close()

	var Product models.Product
	var query string = "%" + params["query"] + "%"
	Product.Name = query
	Product.Description = query

	wg := &sync.WaitGroup{}

	productChan := make(chan []byte, 1)

	wg.Add(1)
	go models.Product.ProductsBySearch(Product, db, productChan, wg)
	wg.Wait()

	close(productChan)

	var Products []models.Product

	if err := json.Unmarshal(<-productChan, &Products); err != nil {
		middleware.SendError(err, res)
		return
	}

	Response := make(map[string]interface{}, 1)
	Response["Products"] = Products

	if err := json.NewEncoder(res).Encode(Response); err != nil {
		middleware.SendError(err, res)
		return
	}
}
