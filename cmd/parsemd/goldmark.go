package main

import (
	"bytes"

	"github.com/sixLUKEY/parsemd.git/internal/plugins"
)

func convert(fileData []byte) (string, error) {
	// Create plugin manager and register plugins
	pluginManager := plugins.NewPluginManager()
	pluginManager.Register(plugins.NewTodoPlugin())

	// Create goldmark instance with plugin extensions
	md := pluginManager.CreateGoldmark()

	// Convert markdown to HTML
	var buf bytes.Buffer
	if err := md.Convert(fileData, &buf); err != nil {
		return "", err
	}

	// Process HTML through plugins
	html := pluginManager.ProcessHTML(buf.String())

	return html, nil
}
