package dao

type ServicesHistory struct {
	BaseModel
	ServiceID      uint            `gorm:"service_id" json:"service_id"`
	ContractID     uint            `gorm:"contract_id" json:"contract_id"`
	IconURL        string          `gorm:"icon_url" json:"icon_url"`
	Price          float64         `gorm:"price" json:"price"`
	IsEnable       bool            `gorm:"is_enable" json:"is_enable"`
	ServiceDemands []ServiceDemand `gorm:"foreignKey:ServiceHistoryID" json:"service_demands"`
	ServiceName    string          `gorm:"service_name" json:"service_name"`
}

func (ServicesHistory) TableName() string {
	return "services_history"
}
