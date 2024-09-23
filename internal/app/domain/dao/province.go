package dao

type Province struct {
	BaseModel
	ProvinceName string `gorm:"province_name" json:"province_name"`
	ProvinceType string `gorm:"province_type" json:"province_type"`
}

func (Province) TableName() string {
	return "province"
}
