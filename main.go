package main

import (
	"embed"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/pazifical/onyx/internal/onyx"
	"github.com/pazifical/onyx/logging"
)

var markdownDirectory = "testdata"

//go:embed frontend/dist
var frontendFS embed.FS

func init() {
	envMarkdownDirectory := os.Getenv("ONYX_MARKDOWN_DIRECTORY")
	if envMarkdownDirectory != "" {
		markdownDirectory = envMarkdownDirectory
	}

	fi, err := os.Stat(markdownDirectory)
	if errors.Is(err, os.ErrNotExist) {
		logging.Info(fmt.Sprintf("markdown directory does not exists: %s", markdownDirectory))
		err = os.MkdirAll(markdownDirectory, 0755)
		if err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		log.Fatal(err)
	}

	if !fi.IsDir() {
		log.Fatalf("given markdown directory is not a directory: %s", markdownDirectory)
	}
}

func main() {
	config := onyx.Config{
		Port:              8080,
		MarkdownDirectory: markdownDirectory,
	}

	server := onyx.NewServer(config, frontendFS)
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
