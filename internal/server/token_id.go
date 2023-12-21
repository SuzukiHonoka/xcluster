package server

import "xcluster/internal/database"

type TokenID uint

func (id TokenID) GetToken() (*Token, error) {
	var token Token
	if err := database.DB.First(&token, "server_token_id", id).Error; err != nil {
		return nil, err
	}
	return &token, nil
}

func (id TokenID) DeleteToken() error {
	return database.DB.Delete(&Token{}, "server_token_id", id).Error
}
