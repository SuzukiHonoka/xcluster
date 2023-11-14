package user

type Group uint

const (
	GroupAdmin Group = iota
	GroupUser
	GroupBanned
)
