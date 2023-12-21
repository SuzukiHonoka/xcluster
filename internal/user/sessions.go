package user

import "xcluster/internal/database"

type Sessions []*Session

func GetSessions() (Sessions, error) {
	var sessions Sessions
	if err := database.DB.Find(sessions).Error; err != nil {
		return nil, err
	}
	return sessions, nil
}

// Invalidate removes the session stored in redis
func (s Sessions) Invalidate() error {
	var err error
	for _, session := range s {
		if err = session.Delete(); err != nil {
			return err
		}
	}
	return nil
}

// todo: auto clean expired session at a random time, expect to be a daily cron job
