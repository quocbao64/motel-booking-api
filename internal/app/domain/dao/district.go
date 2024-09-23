package dao

type District struct {
	BaseModel
	DistrictName string   `gorm:"district_name" json:"district_name"`
	DistrictType string   `gorm:"district_type" json:"district_type"`
	ProvinceID   uint     `gorm:"province_id" json:"province_id"`
	Province     Province `gorm:"foreignKey:ProvinceID" json:"-"`
}

func (District) TableName() string {
	return "district"
}
