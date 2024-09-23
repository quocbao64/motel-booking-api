package repository

import (
	"awesomeProject/internal/app/domain/dao"
	"gorm.io/gorm"
)

type ServiceFilter struct {
}

type ServiceRepository interface {
	GetAll() ([]*dao.Service, error)
	GetByID(id int) (*dao.Service, error)
	Create(service *dao.Service) (*dao.Service, error)
	Update(service *dao.Service) (*dao.Service, error)
	Delete(id int) error
}

type ServiceRepositoryImpl struct {
	db *gorm.DB
}

func (repo ServiceRepositoryImpl) GetAll() ([]*dao.Service, error) {
	var services []*dao.Service
	err := repo.db.Find(&services).Error

	if err != nil {
		return nil, err
	}

	return services, nil
}

func (repo ServiceRepositoryImpl) GetByID(id int) (*dao.Service, error) {
	var service *dao.Service
	err := repo.db.First(&service, id).Error

	if err != nil {
		return &dao.Service{}, err
	}

	return service, nil
}

func (repo ServiceRepositoryImpl) Create(service *dao.Service) (*dao.Service, error) {
	err := repo.db.Create(service).Error

	if err != nil {
		return &dao.Service{}, err
	}

	return service, nil
}

func (repo ServiceRepositoryImpl) Update(service *dao.Service) (*dao.Service, error) {
	err := repo.db.Save(service).Error

	if err != nil {
		return &dao.Service{}, err
	}

	return service, nil
}

func (repo ServiceRepositoryImpl) Delete(id int) error {
	err := repo.db.Delete(&dao.Service{}, id).Error

	if err != nil {
		return err
	}

	return nil
}

func ServiceRepositoryInit(db *gorm.DB) *ServiceRepositoryImpl {
	return &ServiceRepositoryImpl{db: db}
}
