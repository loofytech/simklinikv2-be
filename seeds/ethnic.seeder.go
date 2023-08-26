package seeds

import (
	"sim-klinikv2/models"

	"gorm.io/gorm"
)

func CreateEthnicSeeder(db *gorm.DB) {
	var Ethnics = []models.Ethnic{
		{
			EthnicName: "Jawa",
			EthnicSlug: "jawa",
		},
		{
			EthnicName: "Lampung",
			EthnicSlug: "lampung",
		},
		{
			EthnicName: "Sunda",
			EthnicSlug: "sunda",
		},
		{
			EthnicName: "Bali",
			EthnicSlug: "bali",
		},
		{
			EthnicName: "Minangkabau",
			EthnicSlug: "minangkabau",
		},
		{
			EthnicName: "Betawi",
			EthnicSlug: "betawi",
		},
	}

	for e := range Ethnics {
		db.Model(&models.Ethnic{}).Create(&Ethnics[e])
	}
}
