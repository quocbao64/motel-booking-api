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
		&dao.Address{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
		return
	}
}
