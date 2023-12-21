package server

import (
	"xcluster/internal/database"
	"xcluster/internal/user"
)

// DO NOT TRANSFER THESE METHODS TO user.ID as RECEIVER FUNCTIONS, IT'LL CASE CYCLE IMPORT

func GetGroups(id user.ID) (Groups, error) {
	var groups Groups
	if err := database.DB.Find(&groups, "user_id = ?", id).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

func HasGroupID(id user.ID, gid GroupID) (bool, error) {
	groups, err := GetGroups(id)
	if err != nil {
		return false, err
	}
	return groups.HasGroupID(gid), nil
}
