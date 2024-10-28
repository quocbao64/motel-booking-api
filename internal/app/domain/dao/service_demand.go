package dao

type ServiceDemand struct {
	BaseModel
	OldIndicator     int     `gorm:"old_indicator" json:"old_indicator"`
	NewIndicator     int     `gorm:"new_indicator" json:"new_indicator"`
	ServiceHistoryID uint    `gorm:"service_history_id" json:"service_history_id"`
	InvoiceID        uint    `gorm:"invoice_id" json:"invoice_id"`
	Quality          float64 `gorm:"quality" json:"quality"`
	Amount           float64 `gorm:"amount" json:"amount"`
	AtMonth          int     `gorm:"at_month" json:"at_month"`
	AtYear           int     `gorm:"at_year" json:"at_year"`
	ServiceType      int     `gorm:"service_type" json:"service_type"`
	IsEnable         bool    `gorm:"is_enable" json:"is_enable"`
	BasedPrice       float64 `gorm:"based_price" json:"based_price"`
	Name             string  `gorm:"name" json:"name"`
	Description      string  `gorm:"description" json:"description"`
}
