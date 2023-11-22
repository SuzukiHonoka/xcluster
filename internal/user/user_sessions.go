package user

import "xcluster/internal/database"

type Sessions []*Session

func AllSessions() (Sessions, error) {
	var sessions Sessions
	if err := database.DB.Find(sessions).Error; err != nil {
		return nil, err
	}
	return sessions, nil
}

func (s Sessions) Invalidate() error {
	var err error
	for _, session := range s {
		if err = session.Delete(); err != nil {
			return err
		}
	}
	return nil
}
