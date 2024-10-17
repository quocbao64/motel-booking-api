package repository

import (
	"awesomeProject/internal/app/domain/dao"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ContractFilter struct {
	RenterID int
	LessorID int
}

type ContractRepository interface {
	GetAll(filter *ContractFilter) ([]*dao.Contract, error)
	GetByID(id int) (*dao.Contract, error)
	Create(service *dao.Contract) (*dao.Contract, error)
	Update(service *dao.Contract) (*dao.Contract, error)
	Delete(id int) error
}

type ContractRepositoryImpl struct {
	db *gorm.DB
}

func (repo ContractRepositoryImpl) GetAll(filter *ContractFilter) ([]*dao.Contract, error) {
	var services []*dao.Contract
	db := repo.db

	if filter.RenterID != 0 {
		db = db.Where("renter_id = ?", filter.RenterID)
	}

	if filter.LessorID != 0 {
		db = db.Where("lessor_id = ?", filter.LessorID)
	}

	err := db.Preload(clause.Associations).Find(&services).Error

	if err != nil {
		return nil, err
	}

	return services, nil
}

func (repo ContractRepositoryImpl) GetByID(id int) (*dao.Contract, error) {
	var service *dao.Contract
	err := repo.db.Preload(clause.Associations).First(&service, id).Error

	if err != nil {
		return &dao.Contract{}, err
	}

	return service, nil
}

func (repo ContractRepositoryImpl) Create(service *dao.Contract) (*dao.Contract, error) {
	err := repo.db.Create(service).Error

	if err != nil {
		return &dao.Contract{}, err
	}

	return service, nil
}

func (repo ContractRepositoryImpl) Update(service *dao.Contract) (*dao.Contract, error) {
	err := repo.db.Save(service).Error

	if err != nil {
		return &dao.Contract{}, err
	}

	return service, nil
}

func (repo ContractRepositoryImpl) Delete(id int) error {
	err := repo.db.Delete(&dao.Contract{}, id).Error

	if err != nil {
		return err
	}

	return nil
}

func ContractRepositoryInit(db *gorm.DB) *ContractRepositoryImpl {
	return &ContractRepositoryImpl{db: db}
}
