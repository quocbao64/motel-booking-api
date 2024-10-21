package repository

import (
	"awesomeProject/internal/app/domain/dao"
	"gorm.io/gorm"
)

type ServicesHistoryFilter struct {
}

type ServicesHistoryRepository interface {
	GetAll() ([]*dao.ServicesHistory, error)
	GetByID(id int) (*dao.ServicesHistory, error)
	Create(service *dao.ServicesHistory) (*dao.ServicesHistory, error)
	Update(service *dao.ServicesHistory) (*dao.ServicesHistory, error)
	Delete(id int) error
	CreateMultiple(services []dao.ServicesHistory) ([]dao.ServicesHistory, error)
}

type ServicesHistoryRepositoryImpl struct {
	db *gorm.DB
}

func (repo ServicesHistoryRepositoryImpl) GetAll() ([]*dao.ServicesHistory, error) {
	var services []*dao.ServicesHistory
	err := repo.db.Find(&services).Error

	if err != nil {
		return nil, err
	}

	return services, nil
}

func (repo ServicesHistoryRepositoryImpl) GetByID(id int) (*dao.ServicesHistory, error) {
	var service *dao.ServicesHistory
	err := repo.db.First(&service, id).Error

	if err != nil {
		return &dao.ServicesHistory{}, err
	}

	return service, nil
}

func (repo ServicesHistoryRepositoryImpl) Create(service *dao.ServicesHistory) (*dao.ServicesHistory, error) {
	err := repo.db.Create(service).Error

	if err != nil {
		return &dao.ServicesHistory{}, err
	}

	return service, nil
}

func (repo ServicesHistoryRepositoryImpl) Update(service *dao.ServicesHistory) (*dao.ServicesHistory, error) {
	err := repo.db.Save(service).Error

	if err != nil {
		return &dao.ServicesHistory{}, err
	}

	return service, nil
}

func (repo ServicesHistoryRepositoryImpl) Delete(id int) error {
	err := repo.db.Delete(&dao.ServicesHistory{}, id).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo ServicesHistoryRepositoryImpl) CreateMultiple(services []dao.ServicesHistory) ([]dao.ServicesHistory, error) {
	err := repo.db.Create(&services).Error

	if err != nil {
		return nil, err
	}

	return services, nil
}

func ServicesHistoryRepositoryInit(db *gorm.DB) *ServicesHistoryRepositoryImpl {
	return &ServicesHistoryRepositoryImpl{db: db}
}
