package dao

import "gorm.io/datatypes"

type Room struct {
	BaseModel
	Title         string         `gorm:"title" json:"title"`
	AddressID     uint           `gorm:"address_id" json:"address_id"`
	Acreage       int            `gorm:"acreage" json:"acreage"`
	Price         float64        `gorm:"price" json:"price"`
	Description   string         `gorm:"description" json:"description"`
	DateSubmitted string         `gorm:"date_submitted" json:"date_submitted"`
	OwnerID       uint           `gorm:"owner_id" json:"owner_id"`
	MaxPeople     int            `gorm:"max_people" json:"max_people"`
	RoomType      int            `gorm:"room_type" json:"room_type"`
	Deposit       float64        `gorm:"deposit" json:"deposit"`
	Utilities     string         `gorm:"utilities" json:"utilities"`
	Images        datatypes.JSON `gorm:"images" json:"images"`
}

type RoomResponse struct {
	ID            uint            `json:"id"`
	Title         string          `json:"title"`
	AddressID     uint            `json:"-"`
	Address       AddressResponse `json:"address"`
	Acreage       int             `json:"acreage"`
	Price         float64         `json:"price"`
	Description   string          `json:"description"`
	DateSubmitted string          `json:"date_submitted"`
	OwnerID       uint            `json:"owner_id"`
	MaxPeople     int             `json:"max_people"`
	RoomType      int             `json:"room_type"`
	Deposit       float64         `json:"deposit"`
	Utilities     string          `json:"utilities"`
	Images        datatypes.JSON  `json:"images"`
}
