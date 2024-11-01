package migration

import (
	"awesomeProject/internal/app/domain/dao"
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&dao.Users{},
		&dao.Service{},
		&dao.Province{},
		&dao.District{},
		&dao.Ward{},
		&dao.Room{},
		&dao.Address{},
		&dao.Contract{},
		&dao.Invoice{},
		&dao.HashContract{},
		&dao.ServiceDemand{},
		&dao.BookingRequest{},
		&dao.ServicesHistory{},
		&dao.Transaction{},
		&dao.Signature{},
		&dao.BorrowedItem{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
		return
	}
}
