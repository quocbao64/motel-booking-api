package repository

import (
	"awesomeProject/internal/app/domain/dao"
	"gorm.io/gorm"
)

type GeographyRepository interface {
	GetAllDistrict() ([]*dao.District, error)
	GetDistrictByID(id int) (*dao.District, error)
	GetAllProvince() ([]*dao.Province, error)
	GetProvinceByID(id int) (*dao.Province, error)
	GetAllWard() ([]*dao.Ward, error)
	GetWardByID(id int) (*dao.Ward, error)
}

type GeographyRepositoryImpl struct {
	db *gorm.DB
}

func (repo GeographyRepositoryImpl) GetAllDistrict() ([]*dao.District, error) {
	var districts []*dao.District
	err := repo.db.Find(&districts).Error

	if err != nil {
		return nil, err
	}

	return districts, nil
}

func (repo GeographyRepositoryImpl) GetDistrictByID(id int) (*dao.District, error) {
	var district *dao.District
	err := repo.db.First(&district, id).Error

	if err != nil {
		return &dao.District{}, err
	}

	return district, nil
}

func (repo GeographyRepositoryImpl) GetAllProvince() ([]*dao.Province, error) {
	var provinces []*dao.Province
	err := repo.db.Find(&provinces).Error

	if err != nil {
		return nil, err
	}

	return provinces, nil
}

func (repo GeographyRepositoryImpl) GetProvinceByID(id int) (*dao.Province, error) {
	var province *dao.Province
	err := repo.db.First(&province, id).Error

	if err != nil {
		return &dao.Province{}, err
	}

	return province, nil
}

func (repo GeographyRepositoryImpl) GetAllWard() ([]*dao.Ward, error) {
	var wards []*dao.Ward
	err := repo.db.Find(&wards).Error

	if err != nil {
		return nil, err
	}

	return wards, nil
}

func (repo GeographyRepositoryImpl) GetWardByID(id int) (*dao.Ward, error) {
	var ward *dao.Ward
	err := repo.db.First(&ward, id).Error

	if err != nil {
		return &dao.Ward{}, err
	}

	return ward, nil
}

func GeographyRepositoryInit(db *gorm.DB) *GeographyRepositoryImpl {
	return &GeographyRepositoryImpl{db: db}
}
