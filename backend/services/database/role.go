package database

import "notes/backend/models"

func GetRoleID(code string) (uint, error) {
	var role models.Role

	if err := database.Where("code = ?", code).Find(&role).Error; err != nil {
		return 0, err
	}

	return role.ID, nil
}
