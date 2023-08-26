package seeds

import (
	"sim-klinikv2/models"

	"gorm.io/gorm"
)

func CreateMaritalStatusSeeder(db *gorm.DB) {
	var MaritalStatuss = []models.MaritalStatus{
		{
			MaritalName: "Menikah",
			MaritalSlug: "menikah",
		},
		{
			MaritalName: "Belum Menikah",
			MaritalSlug: "belum-menikah",
		},
		{
			MaritalName: "Cerai Mati",
			MaritalSlug: "cerai-mati",
		},
		{
			MaritalName: "Cerai",
			MaritalSlug: "cerai",
		},
		{
			MaritalName: "Menjanda",
			MaritalSlug: "menjanda",
		},
		{
			MaritalName: "Menduda",
			MaritalSlug: "menduda",
		},
	}

	for e := range MaritalStatuss {
		db.Model(&models.MaritalStatus{}).Create(&MaritalStatuss[e])
	}
}
