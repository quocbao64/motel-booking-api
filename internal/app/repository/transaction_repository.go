package repository

import (
	"awesomeProject/internal/app/constant"
	"awesomeProject/internal/app/domain/dao"
	"gorm.io/gorm"
)

type TransactionFilter struct {
	UserID int
}

type TransactionRepository interface {
	GetAll(filter *TransactionFilter) ([]*dao.Transaction, error)
	GetByID(id int) (*dao.Transaction, error)
	Create(service *dao.Transaction) (*dao.Transaction, error)
	Update(service *dao.Transaction) (*dao.Transaction, error)
	Delete(id int) error
}

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

func (repo TransactionRepositoryImpl) GetAll(filter *TransactionFilter) ([]*dao.Transaction, error) {
	var services []*dao.Transaction
	db := repo.db

	if filter.UserID != 0 {
		db = db.Where("user_id = ?", filter.UserID)
	}

	err := db.Find(&services).Error

	if err != nil {
		return nil, err
	}

	return services, nil
}

func (repo TransactionRepositoryImpl) GetByID(id int) (*dao.Transaction, error) {
	var service *dao.Transaction
	err := repo.db.First(&service, id).Error

	if err != nil {
		return &dao.Transaction{}, err
	}

	return service, nil
}

func (repo TransactionRepositoryImpl) Create(service *dao.Transaction) (*dao.Transaction, error) {
	err := repo.db.Transaction(func(tx *gorm.DB) error {
		var user *dao.Users
		err := tx.First(&user, service.UserID).Error
		if err != nil {
			return err
		}

		if service.Status == constant.TRANSACTION_SUCCESS && (service.TransactionType == constant.TRANSACTION_DEPOSIT || service.TransactionType == constant.TRANSACTION_REFUND) {
			service.BalanceBefore = user.Balance
			user.Balance = user.Balance + service.Amount
			service.BalanceAfter = user.Balance
		} else if service.Status == constant.TRANSACTION_SUCCESS && (service.TransactionType == constant.TRANSACTION_WITHDRAW || service.TransactionType == constant.TRANSACTION_PAYMENT) {
			service.BalanceBefore = user.Balance
			user.Balance = user.Balance - service.Amount
			service.BalanceAfter = user.Balance
		}

		if err := tx.Create(service).Error; err != nil {
			return err
		}

		if err := tx.Save(user).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return &dao.Transaction{}, err
	}

	return service, nil
}

func (repo TransactionRepositoryImpl) Update(service *dao.Transaction) (*dao.Transaction, error) {
	err := repo.db.Save(service).Error

	if err != nil {
		return &dao.Transaction{}, err
	}

	return service, nil
}

func (repo TransactionRepositoryImpl) Delete(id int) error {
	err := repo.db.Delete(&dao.Transaction{}, id).Error

	if err != nil {
		return err
	}

	return nil
}

func TransactionRepositoryInit(db *gorm.DB) *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{db: db}
}
