package seeds

import (
	"sim-klinikv2/models"

	"gorm.io/gorm"
)

func CreateRoleSeeder(db *gorm.DB) {
	var roles = []models.Role{
		{
			RoleName: "Admin",
			RoleSlug: "admin",
		},
		{
			RoleName: "IT",
			RoleSlug: "it",
		},
		{
			RoleName: "Dokter",
			RoleSlug: "dokter",
		},
		{
			RoleName: "Administrasi",
			RoleSlug: "administrasi",
		},
		{
			RoleName: "Perawat",
			RoleSlug: "perawat",
		},
		{
			RoleName: "Apoteker",
			RoleSlug: "apoteker",
		},
		{
			RoleName: "Rujukan",
			RoleSlug: "rujukan",
		},
		{
			RoleName: "Keuangan",
			RoleSlug: "keuangan",
		},
		{
			RoleName: "Kasir",
			RoleSlug: "kasir",
		},
		{
			RoleName: "Radiologi",
			RoleSlug: "radiologi",
		},
		{
			RoleName: "Laboratorium",
			RoleSlug: "laboratorium",
		},
		{
			RoleName: "Fisioterapi",
			RoleSlug: "fisioterapi",
		},
	}

	for h := range roles {
		db.Model(&models.Role{}).Create(&roles[h])
	}
}
