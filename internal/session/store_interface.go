package session

type Store interface {
	Add(session *Session) error
	Get(id ID) (*Session, error)
	Delete(id ID) error
}
