package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: parsemd <command> <file>")
		fmt.Println("Commands:")
		fmt.Println("  parse, p    - Parse and display markdown file")
		fmt.Println("  convert, c  - Convert markdown to HTML")
		os.Exit(1)
	}

	command := os.Args[1]
	filepath := os.Args[2]

	switch command {
	case "parse", "p":
		parse(filepath)
	case "convert", "c":
		convertFile(filepath)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}

func parse(filePath string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(data))
}

func convertFile(filePath string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	html, err := convert(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error converting markdown: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(html)
}
