package seeds

import (
	"sim-klinikv2/models"

	"gorm.io/gorm"
)

func CreatePaymentSeeder(db *gorm.DB) {
	var Payments = []models.Payment{
		{
			PaymentName: "Asuransi",
			PaymentSlug: "asuransi",
		},
		{
			PaymentName: "Bayar Sendiri",
			PaymentSlug: "bayar-sendiri",
		},
	}

	for e := range Payments {
		db.Model(&models.Payment{}).Create(&Payments[e])
	}
}
