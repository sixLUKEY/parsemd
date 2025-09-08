package main

import (
	"fmt"

	"github.com/sixLUKEY/parsemd.git/internal/plugins"
)

func main() {
	// Example 1: Using the utility functions directly
	fmt.Println("=== Todo Utility Functions Demo ===")

	testItems := []string{
		"- [ ] Complete the plugin system",
		"- [x] Implement basic markdown parsing",
		"* [ ] Add more plugin examples",
		"* [x] Set up the project structure",
		"Regular list item",
	}

	for _, item := range testItems {
		fmt.Printf("Item: %q\n", item)
		fmt.Printf("  Is Todo: %v\n", plugins.IsTodoItem(item))
		if plugins.IsTodoItem(item) {
			fmt.Printf("  Is Checked: %v\n", plugins.IsTodoItemChecked(item))
			fmt.Printf("  Todo Text: %q\n", plugins.ExtractTodoText(item))
		}
		fmt.Println()
	}

	// Example 2: Using the TodoProcessor
	fmt.Println("=== Todo Processor Demo ===")

	htmlInput := `<ul>
<li>- [ ] Complete the plugin system</li>
<li>- [x] Implement basic markdown parsing</li>
<li>Regular list item</li>
<li>* [ ] Add more plugin examples</li>
<li>* [x] Set up the project structure</li>
</ul>`

	fmt.Println("Input HTML:")
	fmt.Println(htmlInput)
	fmt.Println()

	processor := plugins.NewTodoProcessor()
	processedHTML := processor.ProcessHTML(htmlInput)

	fmt.Println("Processed HTML:")
	fmt.Println(processedHTML)
	fmt.Println()

	// Example 3: Using the full plugin system
	fmt.Println("=== Full Plugin System Demo ===")

	pluginManager := plugins.NewPluginManager()
	pluginManager.Register(plugins.NewTodoPlugin())

	// Simulate markdown conversion
	markdown := `# Todo List

- [ ] Complete the plugin system
- [x] Implement basic markdown parsing
- [ ] Add more plugin examples
- [x] Set up the project structure

## Regular content

This is regular markdown content.`

	fmt.Println("Markdown Input:")
	fmt.Println(markdown)
	fmt.Println()

	// In a real scenario, this would go through goldmark first
	// For demo purposes, we'll simulate the HTML output
	simulatedHTML := `<h1>Todo List</h1>
<ul>
<li>- [ ] Complete the plugin system</li>
<li>- [x] Implement basic markdown parsing</li>
<li>- [ ] Add more plugin examples</li>
<li>- [x] Set up the project structure</li>
</ul>
<h2>Regular content</h2>
<p>This is regular markdown content.</p>`

	// Process through plugins
	finalHTML := pluginManager.ProcessHTML(simulatedHTML)

	fmt.Println("Final HTML Output:")
	fmt.Println(finalHTML)
}

// Note: These functions would need to be exported in the plugins package
// For this example, we're assuming they are exported as:
// - IsTodoItem
// - IsTodoItemChecked
// - ExtractTodoText
