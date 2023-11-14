package server

type Server struct {
	ID    uint   `gorm:"primaryKey;autoIncrement;unique"`
	Name  string `gorm:"type:varchar(20);not null"`
	Addr  string `gorm:"not null"`  // eg: 127.0.0.1:9999
	Group Group  `gorm:"default:0"` // 0 -> default group
}
