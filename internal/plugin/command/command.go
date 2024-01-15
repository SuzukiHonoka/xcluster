package command

import (
	"xcluster/internal/plugin"
)

var PluginCommand = &command{}

func init() {
	plugin.Store.Register(PluginCommand)
}

type command struct {
	plugin.Plugin
}

func (p *command) Name() string {
	return "command"
}

func (p *command) Describe() string {
	return "Command runner"
}

func (p *command) Author() string {
	return "starx"
}

func (p *command) Actions() plugin.Actions {
	return plugin.ParseActions(p)
}

func (p *command) DescribeActions() plugin.MappedActionDescription {
	return plugin.MappedActionDescription{
		"exec": "execute a command",
	}
}

func (p *command) DescribeActionPayload() plugin.MappedActionPayload {
	return plugin.MappedActionPayload{
		"exec": plugin.Payloads{
			&plugin.Payload{
				Name:        "serverID",
				Type:        plugin.PayloadTypeString,
				Description: "run the command in specified server",
			},
			&plugin.Payload{
				Name:        "command",
				Type:        plugin.PayloadTypeString,
				Description: "command to run",
			},
			&plugin.Payload{
				Name:        "privileged",
				Type:        plugin.PayloadTypeBool,
				Description: "whether run the command in privilege mode",
			},
		},
	}
}

func (p *command) PerformAction(action string, payload plugin.MappedPayload) (interface{}, error) {
	return nil, nil
}
