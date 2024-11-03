package repository

import (
	"awesomeProject/internal/app/domain/dao"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BorrowedItemFilter struct {
	IDs []int
}

type BorrowedItemRepository interface {
	GetAll(filter *BorrowedItemFilter) ([]dao.BorrowedItem, error)
	GetByID(id int) (*dao.BorrowedItem, error)
	Create(service *dao.BorrowedItem) (*dao.BorrowedItem, error)
	Update(service *dao.BorrowedItem) (*dao.BorrowedItem, error)
	Delete(id int) error
	CreateDamagedItems(items []*dao.DamagedItem) error
	GetAllDamagedItems(contractID uint) ([]dao.DamagedItem, error)
	DeleteDamagedItems(contractID uint) error
}

type BorrowedItemRepositoryImpl struct {
	db *gorm.DB
}

func (repo BorrowedItemRepositoryImpl) GetAll(filter *BorrowedItemFilter) ([]dao.BorrowedItem, error) {
	var services []dao.BorrowedItem
	db := repo.db.Preload(clause.Associations)

	if len(filter.IDs) > 0 {
		db = db.Where("id IN (?)", filter.IDs)
	}

	err := db.Find(&services).Error

	if err != nil {
		return nil, err
	}

	return services, nil
}

func (repo BorrowedItemRepositoryImpl) GetByID(id int) (*dao.BorrowedItem, error) {
	var service *dao.BorrowedItem
	err := repo.db.First(&service, id).Error

	if err != nil {
		return &dao.BorrowedItem{}, err
	}

	return service, nil
}

func (repo BorrowedItemRepositoryImpl) Create(service *dao.BorrowedItem) (*dao.BorrowedItem, error) {
	err := repo.db.Create(service).Error

	if err != nil {
		return &dao.BorrowedItem{}, err
	}

	return service, nil
}

func (repo BorrowedItemRepositoryImpl) Update(service *dao.BorrowedItem) (*dao.BorrowedItem, error) {
	err := repo.db.Save(service).Error

	if err != nil {
		return &dao.BorrowedItem{}, err
	}

	return service, nil
}

func (repo BorrowedItemRepositoryImpl) Delete(id int) error {
	err := repo.db.Delete(&dao.BorrowedItem{}, id).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo BorrowedItemRepositoryImpl) CreateDamagedItems(items []*dao.DamagedItem) error {
	err := repo.db.Save(items).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo BorrowedItemRepositoryImpl) GetAllDamagedItems(contractID uint) ([]dao.DamagedItem, error) {
	var items []dao.DamagedItem
	err := repo.db.Where("contract_id = ?", contractID).Find(&items).Error

	if err != nil {
		return nil, err
	}

	return items, nil
}

func (repo BorrowedItemRepositoryImpl) DeleteDamagedItems(contractID uint) error {
	err := repo.db.Where("contract_id = ?", contractID).Delete(&dao.DamagedItem{}).Error

	if err != nil {
		return err
	}

	return nil
}

func BorrowedItemRepositoryInit(db *gorm.DB) *BorrowedItemRepositoryImpl {
	return &BorrowedItemRepositoryImpl{db: db}
}
