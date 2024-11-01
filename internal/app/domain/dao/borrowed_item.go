package dao

type BorrowedItem struct {
	BaseModel
	Name             string  `gorm:"name" json:"name"`
	Price            float64 `gorm:"price" json:"price"`
	RoomID           uint    `gorm:"room_id" json:"room_id"`
	ContractID       *uint   `gorm:"contract_id" json:"-"`
	BookingRequestID *uint   `gorm:"booking_request_id" json:"-"`
}
