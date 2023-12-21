package plugin

type Plugin interface {
	MetaData
	Actions() []string
	DescribeActionPayload() map[string]string
	DescribeActions() map[string]string
	PerformAction(action string, payload map[string]string) (bool, error)
}

type MetaData interface {
	Name() string
	String() string
	Describe() string
	Author() string
}
