package dao

type Transaction struct {
	BaseModel
	UserID          uint    `gorm:"column:user_id" json:"user_id"`
	TransactionType int     `gorm:"column:transaction_type" json:"transaction_type"`
	Status          int     `gorm:"column:status" json:"status"`
	Amount          float64 `gorm:"column:amount" json:"amount"`
	BalanceBefore   float64 `gorm:"column:balance_before" json:"balance_before"`
	BalanceAfter    float64 `gorm:"column:balance_after" json:"balance_after"`
	Description     string  `gorm:"column:description" json:"description"`
	PaymentMethod   int     `gorm:"column:payment_method" json:"payment_method"`
	TransactionNo   string  `gorm:"column:transaction_no" json:"transaction_no"`
}
