package user

type GroupID uint

const (
	GroupIDBanned GroupID = iota
	GroupIDAdmin
	GroupIDUser
)
