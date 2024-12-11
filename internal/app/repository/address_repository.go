package repository

import (
	"awesomeProject/internal/app/domain/dao"
	"gorm.io/gorm"
)

type AddressRepository interface {
	GetFullAddress(id uint) (*dao.AddressResponse, error)
	Create(address *dao.Address) (*dao.Address, error)
	Update(address *dao.Address) (*dao.Address, error)
	GetByID(id uint) (*dao.Address, error)
}

type AddressRepositoryImpl struct {
	db *gorm.DB
}

func AddressRepositoryInit(db *gorm.DB) *AddressRepositoryImpl {
	return &AddressRepositoryImpl{db: db}
}

func (repo AddressRepositoryImpl) GetFullAddress(id uint) (*dao.AddressResponse, error) {
	var address dao.AddressResponse
	err := repo.db.Table("address").
		Select("address.id, address.detail, ward.ward_name, district.district_name, province.province_name, "+
			"ward.id as ward_id, district.id as district_id, province.id as province_id").
		Joins("JOIN ward ON address.ward_id = ward.id").
		Joins("JOIN district ON ward.district_id = district.id").
		Joins("JOIN province ON district.province_id = province.id").
		Where("address.id = ?", id).
		Scan(&address).Error

	if err != nil {
		return &dao.AddressResponse{}, err
	}

	return &address, nil
}

func (repo AddressRepositoryImpl) Create(address *dao.Address) (*dao.Address, error) {
	err := repo.db.Create(&address).Error
	if err != nil {
		return &dao.Address{}, err
	}
	return address, nil
}

func (repo AddressRepositoryImpl) Update(address *dao.Address) (*dao.Address, error) {
	err := repo.db.Save(address).Error
	if err != nil {
		return &dao.Address{}, err
	}
	return address, nil
}

func (repo AddressRepositoryImpl) GetByID(id uint) (*dao.Address, error) {
	var address dao.Address
	err := repo.db.First(&address, id).Error
	if err != nil {
		return &dao.Address{}, err
	}
	return &address, nil
}
