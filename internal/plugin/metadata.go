package plugin

type MetaData interface {
	Name() string
	String() string
	Describe() string
	Author() string
}
