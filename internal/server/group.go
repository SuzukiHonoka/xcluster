package server

import (
	"xcluster/internal/database"
	"xcluster/internal/user"
)

type Group struct {
	ID     GroupID   `gorm:"column:server_group_id;type:int unsigned;primaryKey;autoIncrement;unique" json:"id"`
	Name   GroupName `gorm:"column:server_group_name;type:varchar(255);not null;" json:"name"`
	UserID user.ID   `gorm:"column:user_id;type:int unsigned;not null" json:"userID"`
	User   user.User `gorm:"foreignKey:UserID" json:"-"`
}

func NewGroup(id user.ID, name string) (*Group, error) {
	group := &Group{
		Name:   GroupName(name),
		UserID: id,
	}
	// add to database
	if err := database.DB.Create(group).Error; err != nil {
		return nil, err
	}
	return group, nil
}

func (g *Group) Delete() error {
	servers, err := GetServers(g.ID)
	if err != nil {
		return err
	}
	if err = servers.Delete(); err != nil {
		return err
	}
	if err = g.ID.DeleteGroup(); err != nil {
		return err
	}
	return nil
}

func (*Group) TableName() string {
	return "servers_group"
}
