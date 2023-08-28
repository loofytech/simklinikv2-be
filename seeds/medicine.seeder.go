package seeds

import (
	"sim-klinikv2/models"

	"gorm.io/gorm"
)

func CreateMedicineSeeder(db *gorm.DB) {
	var Medicines = []models.Medicine{
		{
			MedicineName:  "Sabu",
			MedicineType:  "Pill",
			MedicineHPP:   10,
			MedicineHNA:   10,
			MedicineStock: 30,
		},
		{
			MedicineName:  "Ganja",
			MedicineType:  "Pill",
			MedicineHPP:   10,
			MedicineHNA:   10,
			MedicineStock: 30,
		},
	}

	for e := range Medicines {
		db.Model(&models.Medicine{}).Create(&Medicines[e])
	}
}
