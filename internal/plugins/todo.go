package plugins

import (
	"strings"

	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
)

// TodoPlugin processes markdown todo items and converts them to HTML
type TodoPlugin struct{}

// NewTodoPlugin creates a new todo plugin
func NewTodoPlugin() *TodoPlugin {
	return &TodoPlugin{}
}

// Name returns the plugin name
func (tp *TodoPlugin) Name() string {
	return "todo"
}

// ExtendParser adds parser extensions for todo items
func (tp *TodoPlugin) ExtendParser(p parser.Parser) {
	// For now, we'll use the default parser
	// In a full implementation, you'd add custom parsing logic here
}

// ExtendRenderer adds renderer extensions for todo items
func (tp *TodoPlugin) ExtendRenderer(r renderer.Renderer) {
	// For now, we'll use the default renderer
	// In a full implementation, you'd add custom rendering logic here
}

// ProcessHTML processes HTML to transform todo items
func (tp *TodoPlugin) ProcessHTML(html string) string {
	processor := NewTodoProcessor()
	return processor.ProcessHTML(html)
}

// IsTodoItem checks if a text string represents a todo item
func IsTodoItem(text string) bool {
	trimmed := strings.TrimSpace(text)
	return strings.HasPrefix(trimmed, "- [ ]") ||
		strings.HasPrefix(trimmed, "- [x]") ||
		strings.HasPrefix(trimmed, "* [ ]") ||
		strings.HasPrefix(trimmed, "* [x]")
}

// IsTodoItemChecked checks if a todo item is checked
func IsTodoItemChecked(text string) bool {
	return strings.Contains(text, "[x]")
}

// ExtractTodoText extracts the text content from a todo item
func ExtractTodoText(text string) string {
	trimmed := strings.TrimSpace(text)
	// Remove the todo markers
	trimmed = strings.TrimPrefix(trimmed, "- [ ]")
	trimmed = strings.TrimPrefix(trimmed, "- [x]")
	trimmed = strings.TrimPrefix(trimmed, "* [ ]")
	trimmed = strings.TrimPrefix(trimmed, "* [x]")
	return strings.TrimSpace(trimmed)
}

// Internal functions for use within the package
func isTodoItem(text string) bool {
	return IsTodoItem(text)
}

func isTodoItemChecked(text string) bool {
	return IsTodoItemChecked(text)
}

func extractTodoText(text string) string {
	return ExtractTodoText(text)
}
