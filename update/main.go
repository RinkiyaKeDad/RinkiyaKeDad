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

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://dev.to/feed/rinkiyakedad")
	if err != nil {
		log.Fatalf("error getting feed: %v", err)
	}
	// Get the freshest item
	rssItem := feed
	fmt.Println(rssItem.FeedType)
	// Get blog content here and add it to data
	//blog := "Article should be here from dev.to"
	data := fmt.Sprintf("%s\n\n%s\n", stringContent, rssItem)

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
