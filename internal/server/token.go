package server

import (
	"errors"
	"fmt"
	"time"
	"xcluster/internal/database"
	"xcluster/pkg/random"
)

const TokenLength = 16

// todo: cache store

type Token struct {
	ID             TokenID       `gorm:"column:server_token_id;type:int unsigned;primaryKey;autoIncrement;unique" json:"id"`
	TokenRaw       TokenRaw      `gorm:"column:server_token;type:varchar(128)" json:"token"`                                     // optional: hide
	Capacity       TokenCapacity `gorm:"column:server_token_capacity;type:smallint unsigned;not null;default:1" json:"capacity"` // max=65535
	GroupID        GroupID       `gorm:"column:server_group_id;type:int unsigned" json:"groupID"`
	Group          Group         `gorm:"foreignKey:GroupID" json:"-"`
	CreationTime   time.Time     `gorm:"type:datetime" json:"creationTime"`
	ExpirationTime time.Time     `gorm:"type:datetime" json:"expirationTime"`
}

func NewToken(gid GroupID, capacity TokenCapacity, duration time.Duration) (*Token, error) {
	now := time.Now()
	randString := random.String(TokenLength)
	token := &Token{
		TokenRaw:       TokenRaw(randString),
		GroupID:        gid,
		Capacity:       capacity,
		CreationTime:   now,
		ExpirationTime: now.Add(duration),
	}
	// add to database
	if err := database.DB.Create(token).Error; err != nil {
		return nil, err
	}
	return token, nil
}

func (t *Token) Validate() (bool, error) {
	if t.ExpirationTime.Sub(time.Now()) < 0 {
		return false, fmt.Errorf("token expired at %s", t.ExpirationTime)
	}
	if t.Capacity == 0 {
		return false, errors.New("token has zero capacity")
	}
	return true, nil
}

func (t *Token) UseCapacity(delta uint16) (TokenCapacity, error) {
	capacity := t.Capacity - TokenCapacity(delta)
	if err := database.DB.Model(t).Update("server_token_capacity", capacity).Error; err != nil {
		return 0, err
	}
	return capacity, nil
}

func (t *Token) SetCapacity(delta uint16) error {
	return database.DB.Model(t).Update("server_token_capacity", delta).Error
}

func (*Token) TableName() string {
	return "server_token"
}
