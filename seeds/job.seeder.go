package seeds

import (
	"sim-klinikv2/models"

	"gorm.io/gorm"
)

func CreateJobSeeder(db *gorm.DB) {
	var Jobs = []models.Job{
		{
			JobName: "Swasta",
			JobSlug: "swasta",
		},
		{
			JobName: "Pegawai Negeri Sipil",
			JobSlug: "pns",
		},
		{
			JobName: "Polisi",
			JobSlug: "polisi",
		},
		{
			JobName: "Petani",
			JobSlug: "petani",
		},
		{
			JobName: "Ibu Rumah Tangga",
			JobSlug: "irt",
		},
		{
			JobName: "Dokter",
			JobSlug: "dokter",
		},
		{
			JobName: "Tenaga Kesehatan",
			JobSlug: "nakes",
		},
		{
			JobName: "Mahasiswa",
			JobSlug: "mahasiswa",
		},
		{
			JobName: "Dosen",
			JobSlug: "dosen",
		},
		{
			JobName: "Guru",
			JobSlug: "guru",
		},
		{
			JobName: "Pilot",
			JobSlug: "pilot",
		},
		{
			JobName: "Apoteker",
			JobSlug: "apoteker",
		},
	}

	for e := range Jobs {
		db.Model(&models.Job{}).Create(&Jobs[e])
	}
}
