package dao

type Service struct {
	BaseModel
	Name           string          `gorm:"name" json:"name"`
	IconURL        string          `gorm:"icon_url" json:"icon_url"`
	Price          float64         `gorm:"price" json:"price"`
	Description    string          `gorm:"description" json:"description"`
	IsEnable       bool            `gorm:"is_enable" json:"is_enable"`
	ServiceDemands []ServiceDemand `gorm:"foreignKey:ServiceID" json:"service_demands"`
}