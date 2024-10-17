package repository

import (
	"awesomeProject/internal/app/domain/dao"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RoomFilter struct {
	Title   string
	PageID  int
	PerPage int
}

type RoomRepository interface {
	GetAll(filter *RoomFilter) ([]*dao.RoomResponse, error)
	GetByID(id int) (*dao.RoomResponse, error)
	Create(user *dao.Room) (*dao.Room, error)
	Update(user *dao.Room) (*dao.Room, error)
	Delete(id int) error
	CreateRoomService(roomID uint, serviceID uint) error
}

type RoomRepositoryImpl struct {
	db *gorm.DB
}

func (repo RoomRepositoryImpl) GetAll(filter *RoomFilter) ([]*dao.RoomResponse, error) {
	var rooms []*dao.Room

	db := repo.db

	if filter.Title != "" {
		db = db.Where("title LIKE ?", "%"+filter.Title+"%")
	}

	if filter.PageID != 0 && filter.PerPage != 0 {
		db = db.Offset((filter.PageID - 1) * filter.PerPage).Limit(filter.PerPage)
	}

	err := db.Preload(clause.Associations).Find(&rooms).Error

	if err != nil {
		return nil, err
	}

	var roomsResponse []*dao.RoomResponse
	for _, room := range rooms {
		roomsResponse = append(roomsResponse, roomToRoomResponse(room, room.Address.ID))
	}
	return roomsResponse, nil
}

func (repo RoomRepositoryImpl) GetByID(id int) (*dao.RoomResponse, error) {
	var room *dao.Room
	err := repo.db.Preload(clause.Associations).First(&room, id).Error

	if err != nil {
		return &dao.RoomResponse{}, err
	}

	return roomToRoomResponse(room, room.Address.ID), nil
}

func (repo RoomRepositoryImpl) Create(room *dao.Room) (*dao.Room, error) {
	err := repo.db.Create(&room).Error

	if err != nil {
		return &dao.Room{}, err
	}

	return room, nil
}

func (repo RoomRepositoryImpl) Update(room *dao.Room) (*dao.Room, error) {
	err := repo.db.Save(&room).Error

	if err != nil {
		return &dao.Room{}, err
	}

	return room, nil
}

func (repo RoomRepositoryImpl) Delete(id int) error {
	err := repo.db.Delete(&dao.Room{}, id).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo RoomRepositoryImpl) CreateRoomService(roomID uint, serviceID uint) error {
	roomService := &dao.RoomService{
		RoomID:    roomID,
		ServiceID: serviceID,
	}

	err := repo.db.Create(&roomService).Error

	if err != nil {
		return err
	}

	return nil
}

func RoomRepositoryInit(db *gorm.DB) *RoomRepositoryImpl {
	return &RoomRepositoryImpl{db: db}
}

func roomToRoomResponse(room *dao.Room, addressID uint) *dao.RoomResponse {
	return &dao.RoomResponse{
		ID:            room.ID,
		Title:         room.Title,
		Acreage:       room.Acreage,
		Price:         room.Price,
		Description:   room.Description,
		DateSubmitted: room.DateSubmitted,
		OwnerID:       room.OwnerID,
		MaxPeople:     room.MaxPeople,
		RoomType:      room.RoomType,
		Deposit:       room.Deposit,
		Services:      room.Services,
		Images:        room.Images,
		AddressID:     addressID,
	}
}
