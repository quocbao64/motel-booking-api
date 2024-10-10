package repository

import (
	"awesomeProject/internal/app/domain/dao"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ServiceDemandFilter struct {
}

type ServicesDemandRepository interface {
	GetAll() ([]*dao.ServiceDemand, error)
	GetByID(id int) (*dao.ServiceDemand, error)
	Create(service *dao.ServiceDemand) (*dao.ServiceDemand, error)
	Update(service *dao.ServiceDemand) (*dao.ServiceDemand, error)
	Delete(id int) error
}

type ServicesDemandRepositoryImpl struct {
	db *gorm.DB
}

func (repo ServicesDemandRepositoryImpl) GetAll() ([]*dao.ServiceDemand, error) {
	var services []*dao.ServiceDemand
	err := repo.db.Preload(clause.Associations).Find(&services).Error

	if err != nil {
		return nil, err
	}

	return services, nil
}

func (repo ServicesDemandRepositoryImpl) GetByID(id int) (*dao.ServiceDemand, error) {
	var service *dao.ServiceDemand
	err := repo.db.First(&service, id).Error

	if err != nil {
		return &dao.ServiceDemand{}, err
	}

	return service, nil
}

func (repo ServicesDemandRepositoryImpl) Create(service *dao.ServiceDemand) (*dao.ServiceDemand, error) {
	err := repo.db.Create(service).Error

	if err != nil {
		return &dao.ServiceDemand{}, err
	}

	return service, nil
}

func (repo ServicesDemandRepositoryImpl) Update(service *dao.ServiceDemand) (*dao.ServiceDemand, error) {
	err := repo.db.Save(service).Error

	if err != nil {
		return &dao.ServiceDemand{}, err
	}

	return service, nil
}

func (repo ServicesDemandRepositoryImpl) Delete(id int) error {
	err := repo.db.Delete(&dao.ServiceDemand{}, id).Error

	if err != nil {
		return err
	}

	return nil
}

func ServicesDemandRepositoryInit(db *gorm.DB) *ServicesDemandRepositoryImpl {
	return &ServicesDemandRepositoryImpl{db: db}
}
