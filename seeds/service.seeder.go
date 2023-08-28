package seeds

import (
	"sim-klinikv2/models"

	"gorm.io/gorm"
)

func CreateServiceSeeder(db *gorm.DB) {
	var Services = []models.Service{
		{
			ServiceName: "Konsultasi",
			ServiceSlug: "konsultasi",
		},
	}
	for e := range Services {
		db.Model(&models.Service{}).Create(&Services[e])
	}
}
