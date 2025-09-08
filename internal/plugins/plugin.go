package plugins

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
)

// Plugin defines the interface for markdown processing plugins
type Plugin interface {
	// Name returns the unique name of the plugin
	Name() string

	// ExtendParser adds parser extensions to the goldmark parser
	ExtendParser(p parser.Parser)

	// ExtendRenderer adds renderer extensions to the goldmark renderer
	ExtendRenderer(r renderer.Renderer)

	// ProcessHTML processes the final HTML output (optional)
	ProcessHTML(html string) string
}

// PluginManager manages the registration and execution of plugins
type PluginManager struct {
	plugins []Plugin
}

// NewPluginManager creates a new plugin manager
func NewPluginManager() *PluginManager {
	return &PluginManager{
		plugins: make([]Plugin, 0),
	}
}

// Register adds a plugin to the manager
func (pm *PluginManager) Register(plugin Plugin) {
	pm.plugins = append(pm.plugins, plugin)
}

// GetPlugins returns all registered plugins
func (pm *PluginManager) GetPlugins() []Plugin {
	return pm.plugins
}

// CreateGoldmark creates a goldmark instance with all registered plugins
func (pm *PluginManager) CreateGoldmark() goldmark.Markdown {
	md := goldmark.New()

	// Apply all plugin extensions
	for _, plugin := range pm.plugins {
		plugin.ExtendParser(md.Parser())
		plugin.ExtendRenderer(md.Renderer())
	}

	return md
}

// ProcessHTML processes HTML through all registered plugins
func (pm *PluginManager) ProcessHTML(html string) string {
	result := html

	// Apply all plugin HTML processors
	for _, plugin := range pm.plugins {
		result = plugin.ProcessHTML(result)
	}

	return result
}
