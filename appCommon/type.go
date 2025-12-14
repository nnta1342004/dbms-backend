package appCommon

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type SimpleItemModel struct {
	ID       string `json:"id"`
	Quantity int    `json:"quantity"`
}

type SimpleItems []SimpleItemModel

func (j *SimpleItemModel) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Failed to unmarshall json items")
	}
	var item SimpleItemModel
	if err := json.Unmarshal(bytes, &item); err != nil {
		return err
	}
	*j = item
	return nil
}

func (j *SimpleItemModel) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

func (j *SimpleItems) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Failed to unmarshall json items")
	}
	var item SimpleItems
	if err := json.Unmarshal(bytes, &item); err != nil {
		return err
	}
	*j = item
	return nil
}

func (j *SimpleItems) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}
