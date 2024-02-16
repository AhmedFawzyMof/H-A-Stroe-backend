package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"sync"
)

type Category struct {
	Name string `json:"name"`
}

func (c Category) GetAllCategories(db *sql.DB, categoryChan chan []byte, wg *sync.WaitGroup) {
	defer wg.Done()
	var Categories []Category

	categories, err := db.Query("SELECT * FROM Categories")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer categories.Close()

	for categories.Next() {
		var Category Category

		if err := categories.Scan(&Category.Name); err != nil {
			fmt.Println(err.Error())
		}

		Categories = append(Categories, Category)
	}

	categoriesBytes, err := json.Marshal(Categories)

	if err != nil {
		fmt.Println(err.Error())
	}

	categoryChan <- categoriesBytes

}
