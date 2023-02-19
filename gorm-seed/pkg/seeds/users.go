package seeds

import (
	"github.com/ilmsg/gorm-seed/internal/models"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, name string, age int) error {
	return db.Create(&models.User{Name: name, Age: age}).Error
}
