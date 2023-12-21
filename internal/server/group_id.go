package server

import (
	"xcluster/internal/database"
)

type GroupID uint

func (id GroupID) GetGroup() (*Group, error) {
	var group Group
	if err := database.DB.First(&group, "server_group_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

func (id GroupID) DeleteGroup() error {
	return database.DB.Delete(&Group{}, "server_group_id = ?", id).Error
}

func (id GroupID) GetServers() (Servers, error) {
	var servers Servers
	if err := database.DB.Find(&servers, "server_group_id = ?", id).Error; err != nil {
		return nil, err
	}
	return servers, nil
}

func (id GroupID) DeleteServers() error {
	return database.DB.Delete(&Server{}, "server_group_id = ?", id).Error
}

func (id GroupID) GetTokens() (Tokens, error) {
	var tokens Tokens
	if err := database.DB.Find(&tokens, "server_group_id = ?", id).Error; err != nil {
		return nil, err
	}
	return tokens, nil
}

func (id GroupID) DeleteTokens() error {
	return database.DB.Delete(&Token{}, "server_group_id = ?", id).Error
}
