package filesystem

import (
	"errors"
	"os"
	"strings"
)

type Directory struct {
	Path        string      `json:"path"`
	Directories []Directory `json:"directories"`
	Filenames   []string    `json:"filenames"`
}

func CreateDirectoryTree(root string) (Directory, error) {
	rootDirectory := Directory{
		Path:        root,
		Directories: make([]Directory, 0),
		Filenames:   make([]string, 0),
	}

	err := traverseDirectory(root, &rootDirectory)
	if err != nil {
		return Directory{}, err
	}

	return rootDirectory, nil
}

func traverseDirectory(path string, dir *Directory) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			subDir := Directory{
				Path:        path + "/" + entry.Name(),
				Directories: make([]Directory, 0),
				Filenames:   make([]string, 0),
			}

			err := traverseDirectory(subDir.Path, &subDir)
			if err != nil {
				return err
			}

			dir.Directories = append(dir.Directories, subDir)
		} else {
			dir.Filenames = append(dir.Filenames, entry.Name())
		}
	}

	return nil
}

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
		} else if strings.HasSuffix(entry.Name(), ".md") {
			directoryContent.Files = append(directoryContent.Files, entry.Name())
		}
	}

	return directoryContent, nil
}

func CreateDirectory(directoryPath string) error {
	err := os.MkdirAll(directoryPath, 0755)
	if errors.Is(err, os.ErrExist) {
		return nil
	} else if err != nil {
		return err
	}
	return nil
}
