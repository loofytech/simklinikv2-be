package seeds

import (
	"sim-klinikv2/models"

	"gorm.io/gorm"
)

func CreateUserSeeder(db *gorm.DB) {
	var Users = []models.User{
		{
			Name:     "Agung Ardiyanto",
			Username: "agungd3v",
			Email:    "agung@gmail.com",
			Password: "12345",
			RoleId:   3,
		},
	}

	for e := range Users {
		db.Model(&models.User{}).Create(&Users[e])
	}
}
