package plugins

import (
	"testing"
)

func TestIsTodoItem(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"- [ ] Complete task", true},
		{"- [x] Completed task", true},
		{"* [ ] Another task", true},
		{"* [x] Another completed task", true},
		{"Regular list item", false},
		{"  - [ ] Indented todo", true},
		{"- [ ]", true},
		{"- [x]", true},
	}

	for _, test := range tests {
		result := IsTodoItem(test.input)
		if result != test.expected {
			t.Errorf("IsTodoItem(%q) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestIsTodoItemChecked(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"- [x] Completed task", true},
		{"* [x] Another completed task", true},
		{"- [ ] Incomplete task", false},
		{"* [ ] Another incomplete task", false},
		{"- [X] Uppercase X", true},
	}

	for _, test := range tests {
		result := IsTodoItemChecked(test.input)
		if result != test.expected {
			t.Errorf("IsTodoItemChecked(%q) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestExtractTodoText(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"- [ ] Complete task", "Complete task"},
		{"- [x] Completed task", "Completed task"},
		{"* [ ] Another task", "Another task"},
		{"* [x] Another completed task", "Another completed task"},
		{"  - [ ] Indented todo", "Indented todo"},
		{"- [ ]", ""},
		{"- [x]", ""},
	}

	for _, test := range tests {
		result := ExtractTodoText(test.input)
		if result != test.expected {
			t.Errorf("ExtractTodoText(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestTodoProcessor(t *testing.T) {
	processor := NewTodoProcessor()

	input := `<ul>
<li>- [ ] Complete the plugin system</li>
<li>- [x] Implement basic markdown parsing</li>
<li>Regular list item</li>
</ul>`

	expected := `<ul>
<div class="todo-item">
<input type="checkbox" disabled>
<span class="todo-text">Complete the plugin system</span>
</div>
<div class="todo-item completed">
<input type="checkbox" checked disabled>
<span class="todo-text">Implement basic markdown parsing</span>
</div>
<li>Regular list item</li>
</ul>`

	result := processor.ProcessHTML(input)
	if result != expected {
		t.Errorf("ProcessHTML() = %q, expected %q", result, expected)
	}
}
