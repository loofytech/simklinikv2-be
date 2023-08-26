package seeds

import (
	"sim-klinikv2/models"

	"gorm.io/gorm"
)

func CreateReligionSeeder(db *gorm.DB) {
	var Religions = []models.Religion{
		{
			ReligionName: "Islam",
			ReligionSlug: "islam",
		},
		{
			ReligionName: "Kristen",
			ReligionSlug: "kristen",
		},
		{
			ReligionName: "Hindu",
			ReligionSlug: "hindu",
		},
		{
			ReligionName: "Budha",
			ReligionSlug: "Budha",
		},
		{
			ReligionName: "Katolik",
			ReligionSlug: "katolik",
		},
		{
			ReligionName: "Konghucu",
			ReligionSlug: "konghucu",
		},
	}

	for e := range Religions {
		db.Model(&models.Religion{}).Create(&Religions[e])
	}
}
