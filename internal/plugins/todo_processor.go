package plugins

import (
	"regexp"
	"strings"
)

// TodoProcessor processes HTML output to transform todo items
type TodoProcessor struct{}

// NewTodoProcessor creates a new todo processor
func NewTodoProcessor() *TodoProcessor {
	return &TodoProcessor{}
}

// ProcessHTML transforms todo items in HTML
func (tp *TodoProcessor) ProcessHTML(html string) string {
	// Regular expression to match todo items in HTML
	// This matches <li> elements that contain todo patterns
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
