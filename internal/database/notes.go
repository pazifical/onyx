package database

import (
	"fmt"
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

func (nr *NoteRepository) Update(note types.Note) (types.Note, error) {
	filePath := filepath.Join(nr.markdownDirectory, note.Path)

	f, err := os.Create(filePath)
	if err != nil {
		return note, err
	}
	defer f.Close()

	_, err = f.WriteString(note.Text)
	if err != nil {
		return note, err
	}

	return note, nil
}

func (nr *NoteRepository) FetchOne(filePath string) (types.Note, error) {
	data, err := os.ReadFile(filepath.Join(nr.markdownDirectory, filePath))
	if err != nil {
		return types.Note{}, err
	}

	return types.Note{
		Path: filePath,
		Text: string(data),
	}, nil
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
			Path: strings.TrimPrefix(filePath, fmt.Sprintf("%s/", rootDirectory)),
			Text: string(data),
		})

		return nil
	})
	if err != nil {
		return notes, err
	}

	return notes, nil
}
