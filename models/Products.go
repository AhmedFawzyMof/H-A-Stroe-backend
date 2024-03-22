package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"sync"
)

type FilterData struct {
	Min_price int    `json:"min_price"`
	Max_price int    `json:"max_price"`
	Category  string `json:"category"`
}

type Product struct {
	Id          int            `json:"id"`
	Tag         string         `json:"tag"`
	Category    string         `json:"category"`
	Name        string         `json:"name"`
	Slug        string         `json:"slug"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	Image       string         `json:"image"`
	Color       sql.NullString `json:"color"`
}

func (p Product) GetAllProduct(db *sql.DB, productChan chan []byte, wg *sync.WaitGroup) {
	defer wg.Done()

	var Products []Product

	productsPre, err := db.Prepare("SELECT Products.tag, Products.category, Products.name, Products.slug, Products.description, Products.price, ProductImages.image FROM Products INNER JOIN ProductImages ON Products.id = ProductImages.product GROUP BY Products.id")

	if err != nil {
		fmt.Println(err.Error())
	}

	products, err := productsPre.Query()

	if err != nil {
		fmt.Println(err.Error())
	}

	defer products.Close()

	for products.Next() {
		var Product Product

		if err := products.Scan(&Product.Tag, &Product.Category, &Product.Name, &Product.Slug, &Product.Description, &Product.Price, &Product.Image); err != nil {
			fmt.Println(err.Error())
		}

		Products = append(Products, Product)
	}

	ProductsBytes, err := json.Marshal(Products)

	if err != nil {
		fmt.Println(err.Error())
	}

	productChan <- ProductsBytes
}

func (p Product) FilteredProducts(db *sql.DB, filter FilterData, productChan chan []byte, wg *sync.WaitGroup) {
	defer wg.Done()

	var Products []Product

	productsPre, err := db.Prepare("SELECT Products.tag, Products.category, Products.name, Products.slug, Products.description, Products.price, ProductImages.image FROM Products INNER JOIN ProductImages ON Products.id = ProductImages.product WHERE Products.category LIKE ? AND Products.price >= ? AND Products.price <= ? GROUP BY Products.id")

	if err != nil {
		fmt.Println(err.Error())
	}

	products, err := productsPre.Query(filter.Category, filter.Min_price, filter.Max_price)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer products.Close()

	for products.Next() {
		var Product Product

		if err := products.Scan(&Product.Tag, &Product.Category, &Product.Name, &Product.Slug, &Product.Description, &Product.Price, &Product.Image); err != nil {
			fmt.Println(err.Error())
		}

		Products = append(Products, Product)
	}

	ProductsBytes, err := json.Marshal(Products)

	if err != nil {
		fmt.Println(err.Error())
	}

	productChan <- ProductsBytes
}

func (p Product) ProductBySlug(db *sql.DB, productChan chan []byte, wg *sync.WaitGroup) {
	defer wg.Done()

	var Products []Product

	productsPre, err := db.Prepare("SELECT Products.id, Products.tag, Products.category, Products.name, Products.slug, Products.description, Products.price, ProductImages.image, ProductImages.color FROM Products LEFT JOIN ProductImages ON Products.id = ProductImages.product WHERE Products.slug = ?")

	if err != nil {
		fmt.Println(err.Error())
	}

	products, err := productsPre.Query(p.Slug)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer products.Close()

	for products.Next() {

		var Product Product

		if err := products.Scan(&Product.Id, &Product.Tag, &Product.Category, &Product.Name, &Product.Slug, &Product.Description, &Product.Price, &Product.Image, &Product.Color); err != nil {
			fmt.Println(err.Error())
		}

		Products = append(Products, Product)
	}

	var product Product
	if len(Products) > 1 {
		for i := range Products {
			if i != len(Products)-1 {
				product.Image += Products[i].Image + ","
				if Products[i].Color.Valid {
					product.Color.String += Products[i].Color.String + ","
				}
			}
			if i == len(Products)-1 {
				product.Id = Products[i].Id
				product.Name = Products[i].Name
				product.Slug = Products[i].Slug
				product.Description = Products[i].Description
				product.Price = Products[i].Price
				product.Tag = Products[i].Tag
				product.Category = Products[i].Category
				if Products[i].Color.Valid {
					product.Color.String += Products[i].Color.String
				}
				product.Image += Products[i].Image
			}
		}
	} else {
		product = Products[0]
	}

	ProductsBytes, err := json.Marshal(product)

	if err != nil {
		fmt.Println(err.Error())
	}

	productChan <- ProductsBytes
}

func (p Product) ProductsByCategorys(db *sql.DB, productChan chan []byte, wg *sync.WaitGroup) {
	defer wg.Done()

	var Products []Product

	productsPre, err := db.Prepare("SELECT Products.tag, Products.category, Products.name, Products.slug, Products.description, Products.price, ProductImages.image FROM Products INNER JOIN ProductImages ON Products.id = ProductImages.product WHERE Products.category = ? GROUP BY Products.id")

	if err != nil {
		fmt.Println(err.Error())
	}

	products, err := productsPre.Query(p.Category)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer products.Close()

	for products.Next() {
		var Product Product

		if err := products.Scan(&Product.Tag, &Product.Category, &Product.Name, &Product.Slug, &Product.Description, &Product.Price, &Product.Image); err != nil {
			fmt.Println(err.Error())
		}

		Products = append(Products, Product)
	}

	ProductsBytes, err := json.Marshal(Products)

	if err != nil {
		fmt.Println(err.Error())
	}

	productChan <- ProductsBytes
}

func (p Product) ProductsByTag(db *sql.DB, productChan chan []byte, wg *sync.WaitGroup) {
	defer wg.Done()

	var Products []Product

	productsPre, err := db.Prepare("SELECT Products.tag, Products.category, Products.name, Products.slug, Products.description, Products.price, ProductImages.image FROM Products INNER JOIN ProductImages ON Products.id = ProductImages.product WHERE Products.tag = ? GROUP BY Products.id")

	if err != nil {
		fmt.Println(err.Error())
	}

	products, err := productsPre.Query(p.Tag)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer products.Close()

	for products.Next() {
		var Product Product

		if err := products.Scan(&Product.Tag, &Product.Category, &Product.Name, &Product.Slug, &Product.Description, &Product.Price, &Product.Image); err != nil {
			fmt.Println(err.Error())
		}

		Products = append(Products, Product)
	}

	ProductsBytes, err := json.Marshal(Products)

	if err != nil {
		fmt.Println(err.Error())
	}

	productChan <- ProductsBytes
}
func (p Product) ProductsBySearch(db *sql.DB, productChan chan []byte, wg *sync.WaitGroup) {
	defer wg.Done()

	var Products []Product

	productsPre, err := db.Prepare("SELECT Products.tag, Products.category, Products.name, Products.slug, Products.description, Products.price, ProductImages.image FROM Products INNER JOIN ProductImages ON Products.id = ProductImages.product WHERE Products.name LIKE ? OR Products.description LIKE ? GROUP BY Products.id")

	if err != nil {
		fmt.Println(err.Error())
	}

	products, err := productsPre.Query(p.Name, p.Description)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer products.Close()

	for products.Next() {
		var Product Product

		if err := products.Scan(&Product.Tag, &Product.Category, &Product.Name, &Product.Slug, &Product.Description, &Product.Price, &Product.Image); err != nil {
			fmt.Println(err.Error())
		}

		Products = append(Products, Product)
	}

	ProductsBytes, err := json.Marshal(Products)

	if err != nil {
		fmt.Println(err.Error())
	}

	productChan <- ProductsBytes
}
