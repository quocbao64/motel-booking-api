package dao

import (
	"gorm.io/datatypes"
)

type Users struct {
	BaseModel
	FullName        string         `gorm:"full_name" json:"full_name"`
	Email           string         `gorm:"email" json:"email"`
	ImgURL          string         `gorm:"img_url" json:"img_url"`
	Password        string         `gorm:"password" json:"-"`
	Phone           string         `gorm:"phone;not null" json:"phone"`
	Role            datatypes.JSON `gorm:"role" json:"role"`
	RefreshToken    string         `gorm:"refresh_token" json:"-"`
	IdentityNumber  string         `gorm:"identity_number" json:"identity_number"`
	Address         Address        `gorm:"foreignKey:UserID" json:"address"`
	RenterContracts []Contract     `gorm:"foreignKey:RenterID" json:"renter_contracts"`
	LessorContracts []Contract     `gorm:"foreignKey:LessorID" json:"lessor_contracts"`
}

type UsersResponse struct {
	ID             uint            `json:"id"`
	FullName       string          `json:"full_name"`
	Email          string          `json:"email"`
	ImgURL         string          `json:"img_url"`
	Phone          string          `json:"phone"`
	Role           datatypes.JSON  `json:"role"`
	RefreshToken   string          `json:"refresh_token"`
	IdentityNumber string          `json:"identity_number"`
	Address        AddressResponse `json:"address"`
	Password       string          `json:"-"`
}