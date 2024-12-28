package database

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/pazifical/onyx/internal/types"
)

type NoteRepository struct {
	markdownDirectory string
}

func NewNoteRepository(markdownDirectory string) NoteRepository {
	return NoteRepository{
		markdownDirectory: markdownDirectory,
	}
}

func (nr *NoteRepository) FetchAll() ([]types.Note, error) {
	return importAllMarkdownFiles(nr.markdownDirectory)
}

func importAllMarkdownFiles(rootDirectory string) ([]types.Note, error) {
	notes := make([]types.Note, 0)

	err := filepath.WalkDir(rootDirectory, func(filePath string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		if !strings.HasSuffix(filePath, ".md") {
			return nil
		}

		data, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}

		notes = append(notes, types.Note{
			Path: filePath,
			Text: string(data),
		})

		return nil
	})
	if err != nil {
		return notes, err
	}

	return notes, nil
}
