package cache

import (
	"encoding/json"
	"os"
	"time"
)

type Cache struct {
	Data map[string]interface{}
	time.Time
}

func (c *Cache) Set(name string) error {
	c.Data["exp"] = c.Time
	json, err := json.Marshal(c.Data)
	if err != nil {
		return err
	}
	if err := os.WriteFile(name, json, 0777); err != nil {
		return err
	}
	return nil
}

func (c *Cache) Has(key string, fileName string) (bool, error) {
	file, err := os.ReadFile(fileName)

	var cache Cache

	if err != nil {
		return false, err
	}

	if err := json.Unmarshal(file, &cache.Data); err != nil {
		return false, err
	}

	_, has := cache.Data[key]

	return has, nil
}
