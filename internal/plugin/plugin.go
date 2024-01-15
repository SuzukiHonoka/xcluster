package plugin

type Plugins []Plugin

type Plugin interface {
	MetaData
	Actions() Actions
	ActionNames() ActionNames
	DescribeActionPayload() MappedActionPayload
	DescribeActions() MappedActionDescription
	PerformAction(action string, payload MappedPayload) (interface{}, error)
}
