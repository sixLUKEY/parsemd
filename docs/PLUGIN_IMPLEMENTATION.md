# Plugin Implementation Guide

This guide explains how to implement the utility functions in the todo plugin and how to create your own plugins.

## Todo Plugin Utility Functions

The todo plugin provides three main utility functions that you can use to process todo items:

### 1. `IsTodoItem(text string) bool`

Checks if a text string represents a todo item.

**Supported formats:**
- `- [ ] Incomplete task`
- `- [x] Completed task`
- `* [ ] Incomplete task`
- `* [x] Completed task`

**Example:**
```go
import "github.com/sixLUKEY/parsemd.git/internal/plugins"

// Check if text is a todo item
isTodo := plugins.IsTodoItem("- [ ] Complete the plugin system")
// Returns: true

isTodo = plugins.IsTodoItem("Regular list item")
// Returns: false
```

### 2. `IsTodoItemChecked(text string) bool`

Checks if a todo item is marked as completed.

**Example:**
```go
isChecked := plugins.IsTodoItemChecked("- [x] Completed task")
// Returns: true

isChecked = plugins.IsTodoItemChecked("- [ ] Incomplete task")
// Returns: false
```

### 3. `ExtractTodoText(text string) string`

Extracts the text content from a todo item, removing the todo markers.

**Example:**
```go
text := plugins.ExtractTodoText("- [ ] Complete the plugin system")
// Returns: "Complete the plugin system"

text = plugins.ExtractTodoText("- [x] Implement basic markdown parsing")
// Returns: "Implement basic markdown parsing"
```

## How the Functions Work

### Implementation Details

```go
// IsTodoItem checks if a text string represents a todo item
func IsTodoItem(text string) bool {
    trimmed := strings.TrimSpace(text)
    return strings.HasPrefix(trimmed, "- [ ]") || 
           strings.HasPrefix(trimmed, "- [x]") ||
           strings.HasPrefix(trimmed, "* [ ]") ||
           strings.HasPrefix(trimmed, "* [x]")
}
```

**How it works:**
1. Trims whitespace from the input text
2. Checks if the text starts with any of the supported todo patterns
3. Returns `true` if it matches, `false` otherwise

```go
// IsTodoItemChecked checks if a todo item is checked
func IsTodoItemChecked(text string) bool {
    return strings.Contains(text, "[x]")
}
```

**How it works:**
1. Simply checks if the text contains `[x]` (checked marker)
2. Returns `true` if found, `false` otherwise

```go
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
```

**How it works:**
1. Trims whitespace from the input text
2. Removes each possible todo marker using `strings.TrimPrefix`
3. Trims whitespace from the result
4. Returns the clean text content

## TodoProcessor Implementation

The `TodoProcessor` uses these utility functions to transform HTML:

```go
func (tp *TodoProcessor) ProcessHTML(html string) string {
    // Regular expression to match todo items in HTML
    todoRegex := regexp.MustCompile(`<li>(- \[[ x]\]|\* \[[ x]\]) (.+?)</li>`)
    
    return todoRegex.ReplaceAllStringFunc(html, func(match string) string {
        // Extract the todo content
        parts := todoRegex.FindStringSubmatch(match)
        if len(parts) != 3 {
            return match // Return original if no match
        }
        
        todoMarker := parts[1]
        todoText := parts[2]
        
        // Check if it's checked
        isChecked := strings.Contains(todoMarker, "[x]")
        
        // Create the todo HTML
        var todoHTML strings.Builder
        todoHTML.WriteString(`<div class="todo-item`)
        if isChecked {
            todoHTML.WriteString(` completed`)
        }
        todoHTML.WriteString(`">`)
        
        todoHTML.WriteString(`<input type="checkbox"`)
        if isChecked {
            todoHTML.WriteString(` checked`)
        }
        todoHTML.WriteString(` disabled>`)
        
        todoHTML.WriteString(`<span class="todo-text">`)
        todoHTML.WriteString(todoText)
        todoHTML.WriteString(`</span>`)
        
        todoHTML.WriteString(`</div>`)
        
        return todoHTML.String()
    })
}
```

**How it works:**
1. Uses a regular expression to find `<li>` elements containing todo patterns
2. For each match, extracts the todo marker and text
3. Uses the utility functions to determine if it's checked
4. Generates HTML with appropriate CSS classes and attributes
5. Returns the transformed HTML

## Creating Your Own Plugin

To create a new plugin, implement the `Plugin` interface:

```go
type Plugin interface {
    Name() string
    ExtendParser(p parser.Parser)
    ExtendRenderer(r renderer.Renderer)
    ProcessHTML(html string) string
}
```

### Example: Custom Plugin

```go
package plugins

import (
    "strings"
    "github.com/yuin/goldmark/parser"
    "github.com/yuin/goldmark/renderer"
)

type CustomPlugin struct{}

func NewCustomPlugin() *CustomPlugin {
    return &CustomPlugin{}
}

func (cp *CustomPlugin) Name() string {
    return "custom"
}

func (cp *CustomPlugin) ExtendParser(p parser.Parser) {
    // Add custom parsing logic here
}

func (cp *CustomPlugin) ExtendRenderer(r renderer.Renderer) {
    // Add custom rendering logic here
}

func (cp *CustomPlugin) ProcessHTML(html string) string {
    // Add custom HTML processing here
    // Example: Transform custom syntax
    return strings.ReplaceAll(html, "::warning::", "<div class='warning'>")
}
```

## Testing Your Plugin

Use the provided test functions to verify your plugin works correctly:

```go
func TestCustomPlugin(t *testing.T) {
    plugin := NewCustomPlugin()
    
    // Test the plugin name
    if plugin.Name() != "custom" {
        t.Errorf("Expected name 'custom', got %s", plugin.Name())
    }
    
    // Test HTML processing
    input := "Some text with ::warning:: content"
    expected := "Some text with <div class='warning'> content"
    result := plugin.ProcessHTML(input)
    
    if result != expected {
        t.Errorf("Expected %q, got %q", expected, result)
    }
}
```

## Integration with Plugin Manager

Register your plugin with the plugin manager:

```go
pluginManager := plugins.NewPluginManager()
pluginManager.Register(plugins.NewTodoPlugin())
pluginManager.Register(plugins.NewCustomPlugin())

// Create goldmark instance with all plugins
md := pluginManager.CreateGoldmark()

// Convert markdown to HTML
var buf bytes.Buffer
md.Convert(markdownData, &buf)

// Process HTML through all plugins
html := pluginManager.ProcessHTML(buf.String())
```

This architecture allows you to:
- Create reusable utility functions
- Process HTML in a pipeline
- Easily add new plugins
- Test individual components
- Maintain clean separation of concerns
