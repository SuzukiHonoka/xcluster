package server

type Servers []*Server

func (ss Servers) Delete() error {
	var err error
	for _, s := range ss {
		if err = s.Delete(); err != nil {
			return err
		}
	}
	return nil
}
