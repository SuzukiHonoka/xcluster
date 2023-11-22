package server

import (
	"xcluster/internal/database"
)

func GetServers(id GroupID) (Servers, error) {
	var servers Servers
	if err := database.DB.Find(&servers, "server_group_id = ?", id).Error; err != nil {
		return nil, err
	}
	return servers, nil
}
