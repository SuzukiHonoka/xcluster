package server

import (
	"github.com/google/uuid"
	"xcluster/internal/database"
)

type Server struct {
	ID       uint    `gorm:"column:server_pk;type:int unsigned;primaryKey;autoIncrement;unique" json:"-"`
	ServerID ID      `gorm:"column:server_id;type:varchar(36);not null" json:"id"`
	Name     Name    `gorm:"column:server_name;type:varchar(20);not null" json:"name"`
	Addr     Addr    `gorm:"column:server_addr;type:varchar(50);not null" json:"addr"` // eg: 127.0.0.1:9999
	Secret   Secret  `gorm:"column:server_secret;type:varchar(64);not null" json:"-"`  // sha256 digest str with len 64
	GroupID  GroupID `gorm:"column:server_group_id;type:int unsigned;not null" json:"groupID"`
	Group    Group   `gorm:"foreignKey:GroupID" json:"-"`
}

func NewServer(name Name, addr Addr, secret Secret, groupID GroupID) (*Server, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	server := &Server{
		ServerID: ID(id.String()),
		Name:     name,
		Addr:     addr,
		Secret:   secret,
		GroupID:  groupID,
	}
	// add to database
	if err = database.DB.Create(server).Error; err != nil {
		return nil, err
	}
	return server, nil
}

func (s *Server) Delete() error {
	return s.ServerID.DeleteServer()
}

func (*Server) TableName() string {
	return "server"
}
