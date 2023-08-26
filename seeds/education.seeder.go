package seeds

import (
	"sim-klinikv2/models"

	"gorm.io/gorm"
)

func CreateEducationSeeder(db *gorm.DB) {
	var Educations = []models.Education{
		{
			EducationName: "SD",
			EducationSlug: "sd",
		},
		{
			EducationName: "SMP",
			EducationSlug: "smp",
		},
		{
			EducationName: "SMA",
			EducationSlug: "sma",
		},
		{
			EducationName: "D1",
			EducationSlug: "d1",
		},
		{
			EducationName: "D3",
			EducationSlug: "d3",
		},
		{
			EducationName: "D4",
			EducationSlug: "d4",
		},
		{
			EducationName: "S1",
			EducationSlug: "s1",
		},
		{
			EducationName: "S2",
			EducationSlug: "s2",
		},
		{
			EducationName: "S3",
			EducationSlug: "S3",
		},
	}

	for e := range Educations {
		db.Model(&models.Education{}).Create(&Educations[e])
	}
}
