package repository

import (
	"awesomeProject/internal/app/domain/dao"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SignatureFilter struct {
	UserID uint
}

type SignatureRepository interface {
	GetAll(filter *SignatureFilter) ([]*dao.Signature, error)
	GetByID(id int) (*dao.Signature, error)
	Create(service *dao.Signature) (*dao.Signature, error)
	Update(service *dao.Signature) (*dao.Signature, error)
	Delete(id int) error
}

type SignatureRepositoryImpl struct {
	db *gorm.DB
}

func (repo SignatureRepositoryImpl) GetAll(filter *SignatureFilter) ([]*dao.Signature, error) {
	var services []*dao.Signature
	db := repo.db.Preload(clause.Associations)

	if filter.UserID != 0 {
		db = db.Where("user_id = ?", filter.UserID)
	}

	err := db.Find(&services).Error

	if err != nil {
		return nil, err
	}

	return services, nil
}

func (repo SignatureRepositoryImpl) GetByID(id int) (*dao.Signature, error) {
	var service *dao.Signature
	err := repo.db.First(&service, id).Error

	if err != nil {
		return &dao.Signature{}, err
	}

	return service, nil
}

func (repo SignatureRepositoryImpl) Create(service *dao.Signature) (*dao.Signature, error) {
	err := repo.db.Create(service).Error

	if err != nil {
		return &dao.Signature{}, err
	}

	return service, nil
}

func (repo SignatureRepositoryImpl) Update(service *dao.Signature) (*dao.Signature, error) {
	err := repo.db.Save(service).Error

	if err != nil {
		return &dao.Signature{}, err
	}

	return service, nil
}

func (repo SignatureRepositoryImpl) Delete(id int) error {
	err := repo.db.Delete(&dao.Signature{}, id).Error

	if err != nil {
		return err
	}

	return nil
}

func SignatureRepositoryInit(db *gorm.DB) *SignatureRepositoryImpl {
	return &SignatureRepositoryImpl{db: db}
}
