package user

import "time"

// SessionDuration defaults to 3 days
var SessionDuration = 3 * 24 * time.Hour

var AllowRegister = true

func setSessionDuration(duration time.Duration) {
	SessionDuration = duration
}

func setAllowRegister(allow bool) {
	AllowRegister = allow
}
