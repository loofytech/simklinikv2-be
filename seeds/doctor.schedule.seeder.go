package seeds

import (
	"sim-klinikv2/models"

	"gorm.io/gorm"
)

func CreateDoctorScheduleSeeder(db *gorm.DB) {
	var DoctorSchedule = []models.DoctorSchedule{
		{
			UserId:        1,
			UnitId:        1,
			Day:           "Senin",
			OpenPractice:  "12:00",
			ClosePractice: "12:10",
		},
	}

	for e := range DoctorSchedule {
		db.Model(&models.DoctorSchedule{}).Create(&DoctorSchedule[e])
	}
}
