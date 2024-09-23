package dao

type Address struct {
	BaseModel
	WardID uint   `gorm:"ward_id" json:"ward_id"`
	Detail string `gorm:"detail" json:"detail"`
	UserID uint   `gorm:"user_id" json:"user_id"`
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
