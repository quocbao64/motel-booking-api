package repository

import (
	"awesomeProject/internal/app/domain/dao"
	"gorm.io/gorm"
)

type StatisticRepository interface {
	StatisticUser() ([]*dao.UserStatistic, error)
}

type StatisticRepositoryImpl struct {
	db *gorm.DB
}

func (repo StatisticRepositoryImpl) StatisticUser() ([]*dao.UserStatistic, error) {
	var userStatistic []*dao.UserStatistic

	db := repo.db.Table("users").
		Select("id as user_id, full_name, created_at," +
			"(select count(*) from contracts where contracts.renter_id = users.id and status in (2, 7, 8)) as count_of_contract_active_renter," +
			"(select count(*) from contracts where contracts.lessor_id = users.id and status in (3, 5, 6)) as count_of_contract_finished_renter," +
			"(select count(*) from contracts where contracts.lessor_id = users.id and status in (2, 7, 8)) as count_of_contract_active_lessor," +
			"(select count(*) from contracts where contracts.renter_id = users.id and status in (3, 5, 6)) as count_of_contract_finished_lessor," +
			"(select count(*) from contracts where contracts.canceled_by = users.id and contracts.renter_id = users.id) as count_of_canceled_contract_renter," +
			"(select count(*) from contracts where contracts.canceled_by = users.id and contracts.lessor_id = users.id) as count_of_canceled_contract_lessor," +
			"(select count(*) from rooms where rooms.owner_id = users.id) as count_of_rooms")

	err := db.Scan(&userStatistic).Error
	if err != nil {
		return nil, err
	}

	return userStatistic, nil
}

func StatisticRepositoryInit(db *gorm.DB) *StatisticRepositoryImpl {
	return &StatisticRepositoryImpl{db: db}
}
