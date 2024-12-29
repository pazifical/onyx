package filesystem

import (
	"os"
)

type DirectoryContent struct {
	Directories []string `json:"directories"`
	Files       []string `json:"files"`
}

func NewDirectoryContent(directory string) (DirectoryContent, error) {
	directoryContent := DirectoryContent{
		Directories: make([]string, 0),
		Files:       make([]string, 0),
	}

	entries, err := os.ReadDir(directory)
	if err != nil {
		return directoryContent, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			directoryContent.Directories = append(directoryContent.Directories, entry.Name())
		} else {
			directoryContent.Files = append(directoryContent.Files, entry.Name())
		}
	}

	return directoryContent, nil
}
