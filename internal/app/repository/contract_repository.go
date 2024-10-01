package repository

import (
	"awesomeProject/internal/app/domain/dao"
	"gorm.io/gorm"
)

type ContractFilter struct {
}

type ContractRepository interface {
	GetAll() ([]*dao.Contract, error)
	GetByID(id int) (*dao.Contract, error)
	Create(service *dao.Contract) (*dao.Contract, error)
	Update(service *dao.Contract) (*dao.Contract, error)
	Delete(id int) error
	GetAllByRenterOrLessorID(renterID int, lessorID int) ([]*dao.Contract, error)
}

type ContractRepositoryImpl struct {
	db *gorm.DB
}

func (repo ContractRepositoryImpl) GetAll() ([]*dao.Contract, error) {
	var services []*dao.Contract
	err := repo.db.Find(&services).Error

	if err != nil {
		return nil, err
	}

	return services, nil
}

func (repo ContractRepositoryImpl) GetByID(id int) (*dao.Contract, error) {
	var service *dao.Contract
	err := repo.db.First(&service, id).Error

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

func (repo ContractRepositoryImpl) GetAllByRenterOrLessorID(renterID int, lessorID int) ([]*dao.Contract, error) {
	var services []*dao.Contract

	if renterID != 0 {
		err := repo.db.Where("renter_id = ?", renterID).Find(&services).Error

		if err != nil {
			return nil, err
		}

		return services, nil
	} else if lessorID != 0 {
		err := repo.db.Where("lessor_id = ?", lessorID).Find(&services).Error

		if err != nil {
			return nil, err
		}

		return services, nil
	}

	return services, nil
}

func ContractRepositoryInit(db *gorm.DB) *ContractRepositoryImpl {
	return &ContractRepositoryImpl{db: db}
}
