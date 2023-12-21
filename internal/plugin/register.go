package plugin

type Register struct {
	Name                      string
	Description               string
	Author                    string
	Actions                   []string
	ActionsDescription        map[string]string
	ActionsPayloadDescription map[string]string
	Plugin                    Plugin
}

func NewRegister(plugin Plugin) *Register {
	return &Register{
		Name:                      plugin.Name(),
		Description:               plugin.Describe(),
		Author:                    plugin.Author(),
		Actions:                   plugin.Actions(),
		ActionsDescription:        plugin.DescribeActions(),
		ActionsPayloadDescription: plugin.DescribeActionPayload(),
		Plugin:                    plugin,
	}
}
