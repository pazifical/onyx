package main

import (
	"embed"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/pazifical/onyx/internal/matrix"
	"github.com/pazifical/onyx/internal/onyx"
	"github.com/pazifical/onyx/logging"
)

var defaultMarkdownDirectory = "markdown"
var defaultPort = 80

var matrixRoomID string
var matrixUsername string
var matrixPassword string

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

	initMatrixConfigFromEnv()
}

func initMatrixConfigFromEnv() {
	roomID := os.Getenv("ONYX_MATRIX_ROOM_ID")
	if roomID == "" {
		logging.Info("Please provide a Matrix room id via env ONYX_MATRIX_ROOM_ID")
		return
	}
	matrixRoomID = roomID

	username := os.Getenv("ONYX_MATRIX_USERNAME")
	if username == "" {
		logging.Info("Please provide a Matrix username via env ONYX_MATRIX_USERNAME")
		return
	}
	matrixUsername = username

	password := os.Getenv("ONYX_MATRIX_PASSWORD")
	if password == "" {
		logging.Info("Please provide a Matrix user password id via env ONYX_MATRIX_PASSWORD")
		return
	}
	matrixPassword = password
}

func main() {
	config := onyx.Config{
		Port:              defaultPort,
		MarkdownDirectory: defaultMarkdownDirectory,
	}

	server := onyx.NewServer(config, frontendFS)

	if matrixPassword != "" && matrixUsername != "" && matrixRoomID != "" {
		matrixService := matrix.NewService(matrix.Credentials{
			Username: matrixUsername,
			Password: matrixPassword,
		}, matrixRoomID)
		server.AddMatrixService(&matrixService)
	}

	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
