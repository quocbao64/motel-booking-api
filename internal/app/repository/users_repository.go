package repository

import (
	"awesomeProject/internal/app/constant"
	"awesomeProject/internal/app/domain/dao"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	GetAll() ([]*dao.UsersResponse, error)
	GetByID(id int) (*dao.UsersResponse, error)
	GetByEmail(email string) (*dao.Users, error)
	GetByPhone(phone string) (*dao.UsersResponse, error)
	Create(user *dao.Users) (*dao.Users, error)
	Update(user *dao.Users) (*dao.Users, error)
	Delete(id int) error
	UpdateBalance(id uint, transType int, balance float64) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (repo UserRepositoryImpl) GetAll() ([]*dao.UsersResponse, error) {
	var users []*dao.Users
	err := repo.db.Find(users).Error

	if err != nil {
		return nil, err
	}

	var usersResponse []*dao.UsersResponse
	for _, user := range users {
		usersResponse = append(usersResponse, userToUserResponse(user))
	}

	return usersResponse, nil
}

func (repo UserRepositoryImpl) GetByID(id int) (*dao.UsersResponse, error) {
	var user *dao.Users
	err := repo.db.First(&user, id).Error

	if err != nil {
		return &dao.UsersResponse{}, err
	}

	return userToUserResponse(user), nil
}

func (repo UserRepositoryImpl) GetByEmail(email string) (*dao.Users, error) {
	var user *dao.Users
	err := repo.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return &dao.Users{}, err
	}

	return user, nil
}

func (repo UserRepositoryImpl) Create(user *dao.Users) (*dao.Users, error) {
	err := repo.db.Create(&user).Error

	if err != nil {
		return &dao.Users{}, err
	}

	return user, nil
}

func (repo UserRepositoryImpl) Update(user *dao.Users) (*dao.Users, error) {
	err := repo.db.Save(&user).Error

	if err != nil {
		return &dao.Users{}, err
	}

	return user, nil
}

func (repo UserRepositoryImpl) Delete(id int) error {
	err := repo.db.Delete(&dao.Users{}, id).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo UserRepositoryImpl) GetByPhone(phone string) (*dao.UsersResponse, error) {
	var user *dao.Users
	err := repo.db.Where("phone = ?", phone).Preload(clause.Associations).First(&user).Error

	if err != nil {
		return &dao.UsersResponse{}, err
	}

	return userToUserResponse(user), nil
}

func (repo UserRepositoryImpl) UpdateBalance(id uint, transType int, balance float64) error {
	var user *dao.Users
	err := repo.db.First(&user, id).Error

	if err != nil {
		return err
	}

	if transType > 0 {
		if transType == constant.TRANSACTION_DEPOSIT || transType == constant.TRANSACTION_REFUND {
			user.Balance += balance
		} else if transType == constant.TRANSACTION_WITHDRAW || transType == constant.TRANSACTION_PAYMENT {
			user.Balance -= balance
		}
	} else {
		user.Balance = balance
	}

	err = repo.db.Save(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func UserRepositoryInit(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func userToUserResponse(user *dao.Users) *dao.UsersResponse {
	return &dao.UsersResponse{
		ID:             user.ID,
		FullName:       user.FullName,
		Email:          user.Email,
		ImgURL:         user.ImgURL,
		Phone:          user.Phone,
		Role:           user.Role,
		RefreshToken:   user.RefreshToken,
		IdentityNumber: user.IdentityNumber,
		Password:       user.Password,
	}
}
