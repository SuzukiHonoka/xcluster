package user

import (
	"time"
	"xcluster/internal/database"
	"xcluster/internal/session"
)

type Session struct {
	ID             uint       `gorm:"column:session_pk;type:int unsigned;primaryKey;autoIncrement;unique" json:"id"`
	SessionID      session.ID `gorm:"column:session_id;type:varchar(50);not null"` // uuid string length=36
	UserID         ID         `gorm:"column:user_id;type:int unsigned"`
	User           User       `gorm:"foreignKey:UserID"`
	CreationTime   time.Time  `gorm:"type:datetime"`
	ExpirationTime time.Time  `gorm:"type:datetime"`
}

func SaveSession(s *session.Session) (*Session, error) {
	userSession := &Session{
		SessionID:      s.ID,
		UserID:         ID(s.Lease.TenantID),
		CreationTime:   s.Lease.CreationTime,
		ExpirationTime: s.Lease.ExpirationTime,
	}
	if err := database.DB.Create(userSession).Error; err != nil {
		return nil, err
	}
	return userSession, nil
}

func (s *Session) Delete() error {
	if s.ExpirationTime.Sub(time.Now()) <= 0 {
		return nil
	}
	return s.SessionID.DeleteSession()
}

func (*Session) TableName() string {
	return "users_session"
}
