package main

import (
	"log"

	"github.com/pazifical/onyx/internal/onyx"
)

var markdownDirectory = "testdata"

func main() {
	config := onyx.Config{
		Port:              8080,
		MarkdownDirectory: markdownDirectory,
	}

	server := onyx.NewServer(config)
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
