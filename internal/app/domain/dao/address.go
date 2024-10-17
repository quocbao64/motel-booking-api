package dao

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Address struct {
	BaseModel
	WardID          uint   `gorm:"ward_id" json:"ward_id"`
	Detail          string `gorm:"detail" json:"detail"`
	AddressableID   uint   `gorm:"addressable_id" json:"addressable_id"`
	AddressableType string `gorm:"addressable_type" json:"addressable_type"`
}

func (Address) TableName() string {
	return "address"
}

type AddressResponse struct {
	ID           uint   `json:"id"`
	Detail       string `json:"detail"`
	WardName     string `json:"ward_name"`
	DistrictName string `json:"district_name"`
	ProvinceName string `json:"province_name"`
}

func (a Address) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *Address) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, a)
}
