package repository

import (
	"awesomeProject/internal/app/domain/dao"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BookingRequestFilter struct {
	RenterID int
	LessorID int
	PageID   int
	PerPage  int
	RoomID   int
}

type BookingRequestRepository interface {
	GetAll(filter *BookingRequestFilter) ([]*dao.BookingRequest, error)
	GetByID(id int) (*dao.BookingRequest, error)
	Create(bookingRequest *dao.BookingRequest) (*dao.BookingRequest, error)
	Update(bookingRequest *dao.BookingRequest) (*dao.BookingRequest, error)
	Delete(id int) error
	GetByRenterOrLessorID(renterID int, lessorID int) ([]*dao.BookingRequest, error)
}

type BookingRequestRepositoryImpl struct {
	db *gorm.DB
}

func (repo BookingRequestRepositoryImpl) GetAll(filter *BookingRequestFilter) ([]*dao.BookingRequest, error) {
	var bookingRequests []*dao.BookingRequest
	db := repo.db.Preload(clause.Associations).Table("booking_requests").Order("request_date")

	if filter.RenterID != 0 {
		db = db.Where("renter_id = ?", filter.RenterID)
	}

	if filter.LessorID != 0 {
		db = db.Where("lessor_id = ?", filter.LessorID)
	}

	if filter.RoomID != 0 {
		db = db.Where("room_id = ?", filter.RoomID)
	}

	if filter.PageID != 0 && filter.PerPage != 0 {
		db = db.Offset((filter.PageID - 1) * filter.PerPage).Limit(filter.PerPage)
	}

	err := db.Find(&bookingRequests).Error

	if err != nil {
		return nil, err
	}

	return bookingRequests, nil
}

func (repo BookingRequestRepositoryImpl) GetByID(id int) (*dao.BookingRequest, error) {
	var bookingRequest *dao.BookingRequest
	err := repo.db.Preload(clause.Associations).First(&bookingRequest, id).Error

	if err != nil {
		return &dao.BookingRequest{}, err
	}

	return bookingRequest, nil
}

func (repo BookingRequestRepositoryImpl) Create(bookingRequest *dao.BookingRequest) (*dao.BookingRequest, error) {
	err := repo.db.Create(bookingRequest).Error

	if err != nil {
		return &dao.BookingRequest{}, err
	}

	return bookingRequest, nil
}

func (repo BookingRequestRepositoryImpl) Update(bookingRequest *dao.BookingRequest) (*dao.BookingRequest, error) {
	err := repo.db.Save(bookingRequest).Error

	if err != nil {
		return &dao.BookingRequest{}, err
	}

	return bookingRequest, nil
}

func (repo BookingRequestRepositoryImpl) Delete(id int) error {
	err := repo.db.Delete(&dao.BookingRequest{}, id).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo BookingRequestRepositoryImpl) GetByRenterOrLessorID(renterID int, lessorID int) ([]*dao.BookingRequest, error) {
	var bookingRequests []*dao.BookingRequest
	if renterID != 0 {
		err := repo.db.Where("renter_id = ?", renterID).Preload("rooms").Find(&bookingRequests).Order("request_date").Error

		if err != nil {
			return nil, err
		}
	} else if lessorID != 0 {
		err := repo.db.Where("lessor_id = ?", lessorID).Preload("rooms").Find(&bookingRequests).Order("request_date").Error

		if err != nil {
			return nil, err
		}
	}

	return bookingRequests, nil
}

func BookingRequestRepositoryInit(db *gorm.DB) *BookingRequestRepositoryImpl {
	return &BookingRequestRepositoryImpl{db: db}
}
