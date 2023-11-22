package server

import (
	"xcluster/internal/database"
	"xcluster/internal/user"
)

func GetGroups(id user.ID) (Groups, error) {
	var groups Groups
	if err := database.DB.Find(&groups, "user_id = ?", id).Error; err != nil {
		return nil, err
	}
	return groups, nil
}
