# parsemd

A extensible markdown parser with plugin support, built with Go and goldmark.

## Features

- **Plugin Architecture**: Extensible plugin system for custom markdown processing
- **Todo Support**: Built-in todo plugin for interactive checkboxes
- **Multiple Commands**: Parse and convert markdown files
- **Future-Ready**: Designed to support plugins in multiple languages (JS/TS, Lua, Go, Rust)

## Usage

```bash
# Parse and display markdown file content
./parsemd parse test/test.md
./parsemd p test/test.md

# Convert markdown to HTML with plugin processing
./parsemd convert test/test.md
./parsemd c test/test.md
```

## Plugin System

The plugin system allows you to extend markdown processing with custom functionality. Plugins implement the `Plugin` interface:

```go
type Plugin interface {
    Name() string
    ExtendParser(p parser.Parser)
    ExtendRenderer(r renderer.Renderer)
}
```

### Example: Todo Plugin

The included todo plugin converts markdown todo items into interactive HTML:

**Input:**
```markdown
- [ ] Complete the plugin system
- [x] Implement basic markdown parsing
```

**Output:**
```html
<div class="todo-item">
    <input type="checkbox" disabled>
    <span class="todo-text">Complete the plugin system</span>
</div>
<div class="todo-item completed">
    <input type="checkbox" checked disabled>
    <span class="todo-text">Implement basic markdown parsing</span>
</div>
```

## Architecture

The project is designed with future extensibility in mind:

- **Plugin Interface**: Clean interface for adding custom processing
- **Plugin Manager**: Centralized registration and management of plugins
- **Modular Design**: Easy to add new plugins without modifying core code
- **Language Agnostic**: Architecture supports future plugins in different languages

## Development

```bash
# Build the project
go build -o parsemd ./cmd/parsemd

# Run tests
go test ./...

# Test with sample markdown
./parsemd convert test/test.md
```

## Future Roadmap

- [ ] Support for external plugin loading (JS/TS, Lua, etc.)
- [ ] Plugin configuration system
- [ ] More built-in plugins (tables, code blocks, etc.)
- [ ] Plugin marketplace/registry
- [ ] Performance optimizations for large documents