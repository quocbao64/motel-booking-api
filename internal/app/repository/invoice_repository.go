package repository

import (
	"awesomeProject/internal/app/domain/dao"
	"gorm.io/gorm"
)

type InvoiceFilter struct {
}

type InvoiceRepository interface {
	GetAll() ([]*dao.Invoice, error)
	GetByID(id int) (*dao.Invoice, error)
	Create(service *dao.Invoice) (*dao.Invoice, error)
	Update(service *dao.Invoice) (*dao.Invoice, error)
	Delete(id int) error
}

type InvoiceRepositoryImpl struct {
	db *gorm.DB
}

func (repo InvoiceRepositoryImpl) GetAll() ([]*dao.Invoice, error) {
	var services []*dao.Invoice
	err := repo.db.Find(&services).Error

	if err != nil {
		return nil, err
	}

	return services, nil
}

func (repo InvoiceRepositoryImpl) GetByID(id int) (*dao.Invoice, error) {
	var service *dao.Invoice
	err := repo.db.First(&service, id).Error

	if err != nil {
		return &dao.Invoice{}, err
	}

	return service, nil
}

func (repo InvoiceRepositoryImpl) Create(service *dao.Invoice) (*dao.Invoice, error) {
	err := repo.db.Create(service).Error

	if err != nil {
		return &dao.Invoice{}, err
	}

	return service, nil
}

func (repo InvoiceRepositoryImpl) Update(service *dao.Invoice) (*dao.Invoice, error) {
	err := repo.db.Save(service).Error

	if err != nil {
		return &dao.Invoice{}, err
	}

	return service, nil
}

func (repo InvoiceRepositoryImpl) Delete(id int) error {
	err := repo.db.Delete(&dao.Invoice{}, id).Error

	if err != nil {
		return err
	}

	return nil
}

func InvoiceRepositoryInit(db *gorm.DB) *InvoiceRepositoryImpl {
	return &InvoiceRepositoryImpl{db: db}
}
