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
