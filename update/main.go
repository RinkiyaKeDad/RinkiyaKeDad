package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func makeReadme(filename string) error {

	// Unwrap Markdown content
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("cannot read file: %v", err)
		return err
	}

	// Make it a string
	stringContent := string(content)

	// Get blog content here and add it to data
	blog := "Article should be here from dev.to"
	data := fmt.Sprintf("%s\n\n%s\n", stringContent, blog)

	// create the file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Bake at n bytes per second until golden brown
	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

func main() {
	makeReadme("../README.md")
}
