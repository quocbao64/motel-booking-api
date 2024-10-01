package dao

import "time"

type Invoice struct {
	BaseModel
	ContractID     uint            `gorm:"contract_id" json:"contract_id"`
	VAT            float64         `gorm:"vat" json:"vat"`
	Amount         float64         `gorm:"amount" json:"amount"`
	PaymentStatus  int             `gorm:"payment_status" json:"payment_status"`
	PaymentMethod  int             `gorm:"payment_method" json:"payment_method"`
	ServiceDemands []ServiceDemand `gorm:"foreignKey:InvoiceID" json:"service_demands"`
	StartDate      time.Time       `gorm:"start_date" json:"start_date"`
	EndDate        time.Time       `gorm:"end_date" json:"end_date"`
	IsEnable       bool            `gorm:"is_enable" json:"is_enable"`
	IsExtend       bool            `gorm:"is_extend" json:"is_extend"`
	Hash           string          `gorm:"hash" json:"hash"`
}
