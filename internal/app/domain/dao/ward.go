package dao

type Ward struct {
	BaseModel
	WardName   string    `gorm:"ward_name" json:"ward_name"`
	WardType   string    `gorm:"ward_type" json:"ward_type"`
	DistrictID uint      `gorm:"district_id" json:"district_id"`
	District   District  `gorm:"foreignKey:DistrictID" json:"-"`
	Addresses  []Address `gorm:"foreignKey:WardID" json:"addresses"`
}

func (Ward) TableName() string {
	return "ward"
}
