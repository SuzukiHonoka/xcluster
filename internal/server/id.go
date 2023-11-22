package server

import "xcluster/internal/database"

type ID uint

func (id ID) GetServer() (*Server, error) {
	var server Server
	if err := database.DB.First(&server, "server_id", id).Error; err != nil {
		return nil, err
	}
	return &server, nil
}

func (id ID) DeleteServer() error {
	return database.DB.Delete(&Server{}, "server_id", id).Error
}
