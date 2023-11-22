package server

type Groups []*Group

func (gs Groups) Delete() error {
	var err error
	for _, g := range gs {
		// bounded servers
		if err = g.Delete(); err != nil {
			return err
		}
	}
	return nil
}
