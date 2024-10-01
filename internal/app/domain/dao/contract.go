package dao

import "time"

type Contract struct {
	BaseModel
	RenterID      uint           `gorm:"renter_id" json:"renter_id"`
	LessorID      uint           `gorm:"lessor_id" json:"lessor_id"`
	RoomID        uint           `gorm:"room_id" json:"room_id"`
	DateRent      time.Time      `gorm:"date_rent" json:"date_rent"`
	DatePay       time.Time      `gorm:"date_pay" json:"date_pay"`
	PayMode       int            `gorm:"pay_mode" json:"pay_mode"`
	Payment       float64        `gorm:"payment" json:"payment"`
	Status        int            `gorm:"status" json:"status"`
	IsEnable      bool           `gorm:"is_enable" json:"is_enable"`
	FilePath      string         `gorm:"file_path" json:"file_path"`
	Invoices      []Invoice      `gorm:"foreignKey:ContractID" json:"invoices"`
	HashContracts []HashContract `gorm:"foreignKey:ContractID" json:"hash_contracts"`
}
