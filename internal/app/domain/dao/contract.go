package dao

import "time"

type Contract struct {
	BaseModel
	RenterID        uint              `gorm:"renter_id" json:"renter_id"`
	LessorID        uint              `gorm:"lessor_id" json:"lessor_id"`
	RoomID          uint              `gorm:"room_id" json:"room_id"`
	MonthlyPrice    float64           `gorm:"monthly_price" json:"monthly_price"`
	CanceledBy      *uint             `gorm:"canceled_by" json:"canceled_by"`
	DateRent        time.Time         `gorm:"date_rent" json:"date_rent"`
	DatePay         time.Time         `gorm:"date_pay" json:"date_pay"`
	PayMode         int               `gorm:"pay_mode" json:"pay_mode"`
	Payment         float64           `gorm:"payment" json:"payment"`
	Status          int               `gorm:"status" json:"status"`
	IsEnable        bool              `gorm:"is_enable" json:"is_enable"`
	FilePath        string            `gorm:"file_path" json:"file_path"`
	IsRenterSigned  bool              `gorm:"is_renter_signed" json:"is_renter_signed"`
	IsLessorSigned  bool              `gorm:"is_lessor_signed" json:"is_lessor_signed"`
	Invoices        []Invoice         `gorm:"foreignKey:ContractID" json:"invoices"`
	HashContracts   []HashContract    `gorm:"foreignKey:ContractID" json:"hash_contracts"`
	ServiceDemands  []ServiceDemand   `gorm:"foreignKey:ContractID" json:"service_demands"`
	ServicesHistory []ServicesHistory `gorm:"foreignKey:ContractID" json:"services_history"`
}

type ContractResponse struct {
	ID              uint              `json:"id"`
	Renter          UsersResponse     `json:"renter"`
	Lessor          UsersResponse     `json:"lessor"`
	Room            RoomResponse      `json:"room"`
	MonthlyPrice    float64           `json:"monthly_price"`
	CanceledBy      *UsersResponse    `json:"canceled_by"`
	DateRent        time.Time         `json:"date_rent"`
	DatePay         time.Time         `json:"date_pay"`
	PayMode         int               `json:"pay_mode"`
	Payment         float64           `json:"payment"`
	Status          int               `json:"status"`
	IsEnable        bool              `json:"is_enable"`
	FilePath        string            `json:"file_path"`
	Invoices        []Invoice         `json:"invoices"`
	ServiceDemand   []ServiceDemand   `json:"service_demands"`
	ServicesHistory []ServicesHistory `json:"services_history"`
}
