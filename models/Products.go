package models

import (
	"database/sql"
	"errors"
)

type FilterData struct {
	Min_price int    `json:"min_price"`
	Max_price int    `json:"max_price"`
	Category  string `json:"category"`
}

type Product struct {
	Id            int            `json:"id"`
	SubCategory   int            `json:"subcategory"`
	Category      int            `json:"category"`
	Name          string         `json:"name"`
	NameAr        string         `json:"nameAr"`
	Description   string         `json:"description"`
	DescriptionAr string         `json:"descriptionAr"`
	Price         float64        `json:"price"`
	Discount      float64        `json:"discount"`
	Image         string         `json:"image"`
	Color         sql.NullString `json:"color"`
}

// func (p Product) GetAllProduct(db *sql.DB, productChan chan []byte, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	var Products []Product

// 	productsPre, err := db.Prepare("SELECT Products.id, Products.name, Products.nameAr, Products.slug, Products.description, Products.descriptionAr, Products.price, Products.discount, ProductImages.image  FROM Products INNER JOIN ProductImages ON Products.id = ProductImages.product GROUP BY Products.id")

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	products, err := productsPre.Query()

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	defer products.Close()

// 	for products.Next() {
// 		var Product Product

// 		if err := products.Scan(&Product.Id, &Product.Name, &Product.NameAr, &Product.Slug, &Product.Description, &Product.DescriptionAr, &Product.Price, &Product.Discount, &Product.Image); err != nil {
// 			fmt.Println(err.Error())
// 		}

// 		Products = append(Products, Product)
// 	}

// 	ProductsBytes, err := json.Marshal(Products)

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	productChan <- ProductsBytes
// }

// func (p Product) FilteredProducts(db *sql.DB, filter FilterData, productChan chan []byte, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	var Products []Product

// 	productsPre, err := db.Prepare("SELECT Products.id, Products.name, Products.nameAr, Products.slug, Products.description, Products.descriptionAr, Products.price, Products.discount, ProductImages.image FROM Products INNER JOIN ProductImages ON Products.id = ProductImages.product WHERE Products.category LIKE ? AND Products.price >= ? AND Products.price <= ? GROUP BY Products.id")

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	products, err := productsPre.Query(filter.Category, filter.Min_price, filter.Max_price)

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	defer products.Close()

// 	for products.Next() {
// 		var Product Product

// 		if err := products.Scan(&Product.Id, &Product.Name, &Product.NameAr, &Product.Slug, &Product.Description, &Product.DescriptionAr, &Product.Price, &Product.Discount, &Product.Image); err != nil {
// 			fmt.Println(err.Error())
// 		}

// 		Products = append(Products, Product)
// 	}

// 	ProductsBytes, err := json.Marshal(Products)

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	productChan <- ProductsBytes
// }

func (p Product) ProductsByCategorys(db *sql.DB, id, limit int) ([]Product, error) {
	var Products []Product
	var oldLimit int = 0

	if limit > 20 {
		oldLimit = limit /2
	}

	const stableLimit = 20

	productsPre, err := db.Prepare("SELECT Products.id, Products.name, Products.nameAr, Products.description, Products.descriptionAr, Products.price, Products.discount, ProductImages.image FROM Products INNER JOIN ProductImages ON Products.id = ProductImages.product WHERE Products.category = ? GROUP BY Products.id LIMIT ?,?")

	if err != nil {
		return nil, errors.New("error while prossing products")
	}

	products, err := productsPre.Query(id, oldLimit, stableLimit)

	if err != nil {
		return nil, errors.New("error while prossing products")
	}

	defer products.Close()

	for products.Next() {
		var Product Product

		if err := products.Scan(&Product.Id, &Product.Name, &Product.NameAr, &Product.Description, &Product.DescriptionAr, &Product.Price, &Product.Discount, &Product.Image); err != nil {
			return nil, errors.New("error while prossing products")
		}

		Product.Image = "http://localhost:5500/assets" + Product.Image
		Products = append(Products, Product)
	}

	return Products, nil
}

// func (p Product) ProductsByTag(db *sql.DB, productChan chan []byte, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	var Products []Product

// 	productsPre, err := db.Prepare("SELECT Products.id, Products.name, Products.nameAr, Products.slug, Products.description, Products.descriptionAr, Products.price, Products.discount, ProductImages.image FROM Products INNER JOIN ProductImages ON Products.id = ProductImages.product WHERE Products.tag = ? GROUP BY Products.id")

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	products, err := productsPre.Query(p.Tag)

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	defer products.Close()

// 	for products.Next() {
// 		var Product Product

// 		if err := products.Scan(&Product.Id, &Product.Name, &Product.NameAr, &Product.Slug, &Product.Description, &Product.DescriptionAr, &Product.Price, &Product.Discount, &Product.Image); err != nil {
// 			fmt.Println(err.Error())
// 		}

// 		Products = append(Products, Product)
// 	}

// 	ProductsBytes, err := json.Marshal(Products)

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	productChan <- ProductsBytes
// }
// func (p Product) ProductsBySearch(db *sql.DB, productChan chan []byte, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	var Products []Product

// 	productsPre, err := db.Prepare("SELECT Products.id, Products.name, Products.nameAr, Products.slug, Products.description, Products.descriptionAr, Products.price, Products.discount, ProductImages.image FROM Products INNER JOIN ProductImages ON Products.id = ProductImages.product WHERE Products.name LIKE ? OR Products.description LIKE ? GROUP BY Products.id")

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	products, err := productsPre.Query(p.Name, p.Description)

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	defer products.Close()

// 	for products.Next() {
// 		var Product Product

// 		if err := products.Scan(&Product.Id, &Product.Name, &Product.NameAr, &Product.Slug, &Product.Description, &Product.DescriptionAr, &Product.Price, &Product.Discount, &Product.Image); err != nil {
// 			fmt.Println(err.Error())
// 		}

// 		Products = append(Products, Product)
// 	}

// 	ProductsBytes, err := json.Marshal(Products)

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	productChan <- ProductsBytes
// }
