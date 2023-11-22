package server

type Server struct {
	ID      ID     `gorm:"column:server_id;type:int unsigned;primaryKey;autoIncrement;unique"`
	Name    Name   `gorm:"column:server_name;type:varchar(20);not null"`
	Addr    Addr   `gorm:"column:server_addr;type:varchar(50);not null"` // eg: 127.0.0.1:9999
	Secret  Secret `gorm:"column:server_secret;type:tinyblob;not null"`
	GroupID ID     `gorm:"column:server_group_id;type:int unsigned;not null"`
	Group   Group  `gorm:"foreignKey:GroupID"`
}

func NewServer(name, addr string, groupID ID) (*Server, error) {
	server := &Server{
		Name:    Name(name),
		Addr:    Addr(addr),
		GroupID: groupID,
	}
	return server, nil
}

func (s *Server) Delete() error {
	return s.ID.DeleteServer()
}

func (*Server) TableName() string {
	return "servers"
}
