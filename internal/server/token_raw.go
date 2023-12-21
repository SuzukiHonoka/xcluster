package server

import (
	"xcluster/internal/database"
)

type TokenRaw string // random string with length 16

func (r TokenRaw) GetToken() (*Token, error) {
	var token Token
	if err := database.DB.First(&token, "server_token = ?", r).Error; err != nil {
		return nil, err
	}
	return &token, nil
}
