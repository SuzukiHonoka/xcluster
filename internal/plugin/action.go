package plugin

type ActionName string
type ActionNames []ActionName
type ActionDescription string
type MappedActionDescription map[ActionName]ActionDescription
type MappedActionPayload map[ActionName]Payloads

type Actions []*Action

type Action struct {
	plugin      Plugin
	Name        ActionName
	Description ActionDescription
	PayloadInfo Payload
}

func (a *Action) Call(payload PayloadValue) {

}

func ParseActions(plugin Plugin) []*Action {
	actions := make(Actions, 0, len(plugin.ActionNames()))
	// todo parse
	return actions
}
