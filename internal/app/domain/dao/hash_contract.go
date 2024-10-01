package dao

type HashContract struct {
	BaseModel
	ContractID uint   `gorm:"contract_id" json:"contract_id"`
	Hash       string `gorm:"hash" json:"hash"`
}
