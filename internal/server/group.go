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
	var err error
	// delete token
	if err = g.ID.DeleteTokens(); err != nil {
		return err
	}
	// cation, delete the server as well
	if err = g.ID.DeleteServers(); err != nil {
		return err
	}
	// actual delete the group
	if err = g.ID.DeleteGroup(); err != nil {
		return err
	}
	return nil
}

func (*Group) TableName() string {
	return "server_group"
}
