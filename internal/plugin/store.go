package plugin

import (
	"fmt"
)

var Store = make(store)

// store stores the plugin ptr, use plugin name or uuid as key
type store map[string]Plugin

func (s store) Register(plugin Plugin) {
	s[plugin.Name()] = plugin
}

func (s store) UnRegister(plugin Plugin) {
	delete(s, plugin.Name())
}

func (s store) GetPlugin(name string) (Plugin, bool) {
	p, ok := s[name]
	return p, ok
}

func (s store) ListPlugins() Plugins {
	plugins := make(Plugins, 0, len(s))
	for _, plugin := range s {
		plugins = append(plugins, plugin)
	}
	return plugins
}

func (s store) ListPluginNames() []string {
	names := make([]string, 0, len(s))
	for _, plugin := range s {
		names = append(names, plugin.Name())
	}
	return names
}

func (s store) Call(name, action string, payload MappedPayload) (interface{}, error) {
	p, ok := s.GetPlugin(name)
	if !ok {
		return nil, fmt.Errorf("plugin with name=%s not found", name)
	}
	return p.PerformAction(action, payload)
}
