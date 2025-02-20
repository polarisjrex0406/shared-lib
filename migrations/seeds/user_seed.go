package seeds

import (
	"github.com/omimic12/shared-lib/entities"
	"github.com/omimic12/shared-lib/utils"
	"gorm.io/gorm"
)

func User(db *gorm.DB) error {
	hashedPassword, err := utils.HashPassword("123456aA@")
	if err != nil {
		return err
	}
	superuser := entities.User{
		Email:     "superuser@mail.com",
		Firstname: "Super",
		Lastname:  "User",
		Password:  hashedPassword,
		Role:      entities.RoleAdmin,
		Status:    entities.StatusIdle,
	}
	if err := db.Where("email = ?", superuser.Email).Attrs(superuser).FirstOrCreate(&superuser).Error; err != nil {
		return err
	}
	return nil
}
