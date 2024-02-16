package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"sync"
)

type Tag struct {
	Name string `json:"name"`
}

func (t Tag) GetAllTags(db *sql.DB, tagChan chan []byte, wg *sync.WaitGroup) {
	defer wg.Done()
	var Tags []Tag

	tags, err := db.Query("SELECT * FROM Tags")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer tags.Close()

	for tags.Next() {
		var Tag Tag

		if err := tags.Scan(&Tag.Name); err != nil {
			fmt.Println(err.Error())
		}

		Tags = append(Tags, Tag)
	}

	tagsBytes, err := json.Marshal(Tags)

	if err != nil {
		fmt.Println(err.Error())
	}

	tagChan <- tagsBytes

}
