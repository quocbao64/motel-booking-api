package repository

import (
	"awesomeProject/internal/app/domain/dao"
	"gorm.io/gorm"
)

type HashContractFilter struct {
}

type HashContractRepository interface {
	GetAll() ([]*dao.HashContract, error)
	GetByID(id int) (*dao.HashContract, error)
	Create(service *dao.HashContract) (*dao.HashContract, error)
	Update(service *dao.HashContract) (*dao.HashContract, error)
	Delete(id int) error
	GetByContractID(contractID int) (*dao.HashContract, error)
}

type HashContractRepositoryImpl struct {
	db *gorm.DB
}

func (repo HashContractRepositoryImpl) GetAll() ([]*dao.HashContract, error) {
	var services []*dao.HashContract
	err := repo.db.Find(&services).Error

	if err != nil {
		return nil, err
	}

	return services, nil
}

func (repo HashContractRepositoryImpl) GetByID(id int) (*dao.HashContract, error) {
	var service *dao.HashContract
	err := repo.db.First(&service, id).Error

	if err != nil {
		return &dao.HashContract{}, err
	}

	return service, nil
}

func (repo HashContractRepositoryImpl) Create(service *dao.HashContract) (*dao.HashContract, error) {
	err := repo.db.Create(service).Error

	if err != nil {
		return &dao.HashContract{}, err
	}

	return service, nil
}

func (repo HashContractRepositoryImpl) Update(service *dao.HashContract) (*dao.HashContract, error) {
	err := repo.db.Save(service).Error

	if err != nil {
		return &dao.HashContract{}, err
	}

	return service, nil
}

func (repo HashContractRepositoryImpl) Delete(id int) error {
	err := repo.db.Delete(&dao.HashContract{}, id).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo HashContractRepositoryImpl) GetByContractID(contractID int) (*dao.HashContract, error) {
	var service *dao.HashContract
	err := repo.db.Where("contract_id = ?", contractID).First(&service).Error

	if err != nil {
		return &dao.HashContract{}, err
	}

	return service, nil
}

func HashContractRepositoryInit(db *gorm.DB) *HashContractRepositoryImpl {
	return &HashContractRepositoryImpl{db: db}
}
