package seeds

import (
	"sim-klinikv2/models"

	"gorm.io/gorm"
)

func CreateRoleSeeder(db *gorm.DB) {
	var roles = []models.Role{
		{
			RoleName: "Role 1",
			RoleSlug: "role_1",
		},
		{
			RoleName: "Role 1",
			RoleSlug: "role_2",
		},
	}

	for h := range roles {
		db.Model(&models.Role{}).Create(&roles[h])
	}
}
