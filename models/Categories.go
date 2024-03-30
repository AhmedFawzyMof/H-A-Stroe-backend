package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"sync"
)

type Category struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	NameAr string `json:"nameAr"`
	Img    string `json:"img"`
}

func (c Category) GetAllCategories(db *sql.DB, categoryChan chan []byte, wg *sync.WaitGroup, img bool) {
	defer wg.Done()
	var Categories []Category

	var sqlstmt string = "SELECT * FROM Categories"
	if !img {
		sqlstmt = "SELECT id, name, nameAr FROM Categories"
	}

	categories, err := db.Query(sqlstmt)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer categories.Close()

	for categories.Next() {
		var Category Category

		if img {
			err := categories.Scan(&Category.Id, &Category.Name, &Category.NameAr, &Category.Img)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			err := categories.Scan(Category.Id, &Category.Name, &Category.NameAr)
			if err != nil {
				fmt.Println(err.Error())
			}
		}

		Categories = append(Categories, Category)
	}

	categoriesBytes, err := json.Marshal(Categories)

	if err != nil {
		fmt.Println(err.Error())
	}

	categoryChan <- categoriesBytes

}
