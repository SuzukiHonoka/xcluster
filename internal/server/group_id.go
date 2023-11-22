package server

import (
	"xcluster/internal/database"
)

type GroupID uint

func (id GroupID) GetGroup() (*Group, error) {
	var group Group
	if err := database.DB.First(&group, "server_group_id", id).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

func (id GroupID) DeleteGroup() error {
	return database.DB.Delete(&Group{}, "server_group_id", id).Error
}
