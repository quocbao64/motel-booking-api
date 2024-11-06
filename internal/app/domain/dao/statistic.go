package dao

import "time"

type UserStatistic struct {
	UserID                        int       `json:"user_id"`
	FullName                      string    `json:"full_name"`
	CreatedAt                     time.Time `json:"created_at"`
	CountOfRooms                  int       `json:"count_of_rooms"`
	CountOfContractActiveRenter   int       `json:"count_of_contract_active_renter"`
	CountOfContractFinishedRenter int       `json:"count_of_contract_finished_renter"`
	CountOfContractActiveLessor   int       `json:"count_of_contract_active_lessor"`
	CountOfContractFinishedLessor int       `json:"count_of_contract_finished_lessor"`
	CountOfCanceledContractRenter int       `json:"count_of_canceled_contract_renter"`
	CountOfCanceledContractLessor int       `json:"count_of_canceled_contract_lessor"`
}

type RoomStatistic struct {
	RoomID                int       `json:"room_id"`
	WardName              string    `json:"ward_name"`
	LessorID              int       `json:"lessor_id"`
	LessorName            string    `json:"lessor_name"`
	CountOfBookingRequest int       `json:"count_of_booking_request"`
	CountOfContract       int       `json:"count_of_contract"`
	CountOfRenewal        int       `json:"count_of_renewal"`
	CreatedAt             time.Time `json:"created_at"`
}
