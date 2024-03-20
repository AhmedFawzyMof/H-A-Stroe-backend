package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"sync"
)

type Category struct {
	Name string `json:"name"`
	Img  string `json:"img"`
}

func (c Category) GetAllCategories(db *sql.DB, categoryChan chan []byte, wg *sync.WaitGroup, img bool) {
	defer wg.Done()
	var Categories []Category

	var sqlstmt string = "SELECT name, img FROM Categories"
	if !img {
		sqlstmt = "SELECT name FROM Categories"
	}

	categories, err := db.Query(sqlstmt)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer categories.Close()

	for categories.Next() {
		var Category Category

		if img {
			err := categories.Scan(&Category.Name, &Category.Img)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			err := categories.Scan(&Category.Name)
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
