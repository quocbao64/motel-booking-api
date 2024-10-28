package dao

type Signature struct {
	BaseModel
	UserID      uint   `gorm:"user_id" json:"user_id"`
	ServiceType string `gorm:"service_type" json:"service_type"`
	CCCDNumber  string `gorm:"cccd_number" json:"cccd_number"`
}
