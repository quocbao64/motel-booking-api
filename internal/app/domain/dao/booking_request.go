package dao

import "time"

type BookingRequest struct {
	BaseModel
	RenterID          uint      `gorm:"renter_id" json:"renter_id"`
	LessorID          uint      `gorm:"lessor_id" json:"lessor_id"`
	RoomID            uint      `gorm:"room_id" json:"room_id"`
	Room              Room      `gorm:"-" json:"room"`
	RequestDate       time.Time `gorm:"request_date" json:"request_date"`
	Status            string    `gorm:"status" json:"status"`
	Note              string    `gorm:"note" json:"note"`
	MessageFromRenter string    `gorm:"message_from_renter" json:"message_from_renter"`
	MessageFromLessor string    `gorm:"message_from_lessor" json:"message_from_lessor"`
	StartDate         time.Time `gorm:"start_date" json:"start_date"`
	RentalDuration    int       `gorm:"rental_duration" json:"rental_duration"`
	ResponseDate      time.Time `gorm:"response_date" json:"response_date"`
	ContractID        uint      `gorm:"contract_id" json:"contract_id"`
}
