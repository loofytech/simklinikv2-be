package seeds

import (
	"sim-klinikv2/models"

	"gorm.io/gorm"
)

func CreateUnitSeeder(db *gorm.DB) {
	var Units = []models.Unit{
		{
			UnitName:  "Pharmacy",
			UnitSlug:  "pharmacy",
			ServiceId: 1,
		},
	}

	for e := range Units {
		db.Model(&models.Unit{}).Create(&Units[e])
	}
}
