package session

type ID string // uuid string

func (id ID) GetSession() (*Session, error) {
	return store.Get(id)
}

func (id ID) DeleteSession() error {
	return store.Delete(id)
}
