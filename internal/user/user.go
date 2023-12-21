package user

import (
	"errors"
	"fmt"
	"xcluster/internal/database"
	"xcluster/pkg/argon2"
)

type User struct {
	ID       ID       `gorm:"column:user_id;type:int unsigned;primaryKey;autoIncrement;unique" json:"id"` // uid 0 is not allowed
	Name     Name     `gorm:"column:user_name;type:varchar(100);unique;not null" json:"name"`             // account name
	Password Password `gorm:"column:user_password;type:varchar(255);not null" json:"-"`                   // argon2 hashed with salt
	Email    Email    `gorm:"column:user_email;type:varchar(100);unique;not null" json:"email"`
	GroupID  GroupID  `gorm:"column:user_group_id;type:int unsigned;default:1" json:"groupID"` // 0 -> baned, 1 -> admin, 2 -> user, ...
}

func NewUser(name, password, email string) (*User, error) {
	hash, err := argon2.NewArgon2(nil).GenerateHashFromString(password)
	if err != nil {
		return nil, err
	}
	user := &User{
		// note that ID is generated by database
		Name:     Name(name),
		Password: Password(hash.String()),
		Email:    Email(email),
		GroupID:  GroupIDAdmin, // admin as default for now
	}
	// add to database
	if err = database.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) IsAdmin() bool {
	return u.GroupID == GroupIDAdmin
}

func (u *User) Update(name, password, email string) error {
	if name == "" && password == "" && email == "" {
		return errors.New("all fields empty")
	}
	val := map[string]interface{}{}
	if name != "" {
		val["user_name"] = name
	}
	if password != "" {
		hash, err := argon2.NewArgon2(nil).GenerateHashFromString(password)
		if err != nil {
			return err
		}
		val["user_password"] = hash.String()
	}
	if email != "" {
		val["user_email"] = email
	}
	// optional: update receiver properties as well
	return database.DB.Model(u).Updates(val).Error
}

func (u *User) Delete() error {
	return u.ID.DeleteUser()
}

func (u *User) String() string {
	return fmt.Sprintf("[user] name=%s, email=%s", u.Name, u.Email)
}

func (*User) TableName() string {
	return "user"
}
