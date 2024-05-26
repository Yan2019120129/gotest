package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
)

type GormStrings []string

func (_GormStrings GormStrings) Value() (driver.Value, error) {
	if _GormStrings == nil {
		return "[]", nil
	}
	return json.Marshal(_GormStrings)
}

func (_GormStrings *GormStrings) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to scan Array value:", value))
	}

	if len(bytes) > 0 {
		return json.Unmarshal(bytes, _GormStrings)
	}
	*_GormStrings = make([]string, 0)
	return nil
}
