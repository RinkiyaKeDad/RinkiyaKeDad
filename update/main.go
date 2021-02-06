package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/mmcdole/gofeed"
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

	// Get rss feed from dev.to
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://dev.to/feed/rinkiyakedad")
	if err != nil {
		log.Fatalf("error getting feed: %v", err)
	}

	// add latest 5 blogs to a string
	var blog string
	for i := 0; i < 5; i++ {
		blogItem := feed.Items[i]
		blog += "**[" + blogItem.Title + "](" + blogItem.Link + ")**<br/>"
	}

	// add blogs to data
	data := fmt.Sprintf("%s\n\n%s\n", stringContent, blog)

	// create the file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// write the file
	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

func main() {
	makeReadme("../README.md")
}
