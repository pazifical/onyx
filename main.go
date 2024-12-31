package main

import (
	"embed"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/pazifical/onyx/internal/onyx"
	"github.com/pazifical/onyx/logging"
)

var defaultMarkdownDirectory = "markdown"
var defaultPort = 80

//go:embed frontend/dist
var frontendFS embed.FS

var isDevMode bool

func init() {
	envPortString := os.Getenv("ONYX_PORT")
	if envPortString != "" {
		envPort, err := strconv.Atoi(envPortString)
		if err == nil {
			defaultPort = envPort
		}
	}

	envMarkdownDirectory := os.Getenv("ONYX_MARKDOWN_DIRECTORY")
	if envMarkdownDirectory != "" {
		defaultMarkdownDirectory = envMarkdownDirectory
	}

	fi, err := os.Stat(defaultMarkdownDirectory)
	if errors.Is(err, os.ErrNotExist) {
		logging.Info(fmt.Sprintf("markdown directory does not exists: %s", defaultMarkdownDirectory))
		err = os.MkdirAll(defaultMarkdownDirectory, 0755)
		if err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		log.Fatal(err)
	}

	if !fi.IsDir() {
		log.Fatalf("given markdown directory is not a directory: %s", defaultMarkdownDirectory)
	}
}

func main() {
	config := onyx.Config{
		Port:              defaultPort,
		MarkdownDirectory: defaultMarkdownDirectory,
	}

	server := onyx.NewServer(config, frontendFS)
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
