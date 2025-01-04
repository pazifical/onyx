package onyx

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/pazifical/onyx/internal"
	"github.com/pazifical/onyx/internal/filesystem"
	"github.com/pazifical/onyx/internal/types"
	"github.com/pazifical/onyx/logging"
)

func (s *Server) GetAllNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := s.noteRepository.FetchAll()
	if err != nil {
		logging.Error(err.Error())
		s.respondWithError(w, http.StatusInternalServerError, internal.OnyxError{
			ErrorMessage: "unable to fetch notes from markdown files",
		})
		return
	}

	err = json.NewEncoder(w).Encode(notes)
	if err != nil {
		logging.Error(err.Error())
		s.respondWithError(w, http.StatusInternalServerError, internal.OnyxError{
			ErrorMessage: "unable to encode notes to JSON",
		})
		return
	}
}

func (s *Server) GetNoteByID(w http.ResponseWriter, r *http.Request) {
	noteID := r.PathValue("id")

	note, err := s.noteRepository.FetchOne(noteID)
	if err != nil {
		logging.Error(err.Error())
		s.respondWithError(w, http.StatusInternalServerError, internal.OnyxError{
			ErrorMessage: fmt.Sprintf("unable to fetch note from markdown file: %s", noteID),
		})
		return
	}

	err = json.NewEncoder(w).Encode(note)
	if err != nil {
		logging.Error(err.Error())
		s.respondWithError(w, http.StatusInternalServerError, internal.OnyxError{
			ErrorMessage: fmt.Sprintf("unable to encode note '%s' to JSON", noteID),
		})
		return
	}
}

func (s *Server) SaveNote(w http.ResponseWriter, r *http.Request) {
	var note types.Note

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		logging.Error(err.Error())
		s.respondWithError(w, http.StatusNotAcceptable, internal.OnyxError{
			ErrorMessage: "unable to parse note",
		})
		return
	}

	note, err = s.noteRepository.Update(note)
	if err != nil {
		logging.Error(err.Error())
		s.respondWithError(w, http.StatusInternalServerError, internal.OnyxError{
			ErrorMessage: "unable to update note markdown file",
		})
		return
	}

	err = json.NewEncoder(w).Encode(note)
	if err != nil {
		logging.Error(err.Error())
		s.respondWithError(w, http.StatusInternalServerError, internal.OnyxError{
			ErrorMessage: "unable to encode updated note to JSON",
		})
		return
	}
}

func (s *Server) GetNoteByFilePath(w http.ResponseWriter, r *http.Request) {
	filePath := r.PathValue("path")

	if filePath == "" {
		s.respondWithError(w, http.StatusNotAcceptable, internal.OnyxError{
			ErrorMessage: "missing filepath to markdown file",
		})
		return
	} else if !strings.HasSuffix(filePath, ".md") {
		s.respondWithError(w, http.StatusNotAcceptable, internal.OnyxError{
			ErrorMessage: "given markdown path does not end with '.md'",
		})
		return
	}

	filePath = strings.ReplaceAll(filePath, "%20", " ")

	note, err := s.noteRepository.FetchOne(filePath)
	if err != nil {
		logging.Error(err.Error())
		s.respondWithError(w, http.StatusInternalServerError, internal.OnyxError{
			ErrorMessage: fmt.Sprintf("unable to fetch note from markdown file: %s", filePath),
		})
		return
	}

	err = json.NewEncoder(w).Encode(note)
	if err != nil {
		logging.Error(err.Error())
		s.respondWithError(w, http.StatusInternalServerError, internal.OnyxError{
			ErrorMessage: "unable to encode updated note to JSON",
		})
		return
	}
}

func (s *Server) CreateDirectory(w http.ResponseWriter, r *http.Request) {
	directoryPath := r.PathValue("path")
	directoryPath = strings.ReplaceAll(directoryPath, "%20", " ")
	directoryPath = filepath.Join(s.config.MarkdownDirectory, directoryPath)

	err := filesystem.CreateDirectory(directoryPath)
	if err != nil {
		logging.Error(err.Error())
		s.respondWithError(w, http.StatusInternalServerError, internal.OnyxError{
			ErrorMessage: fmt.Sprintf("unable to create directory: %s", directoryPath),
		})
		return
	}

	s.GetDirectoryContent(w, r)
}

func (s *Server) GetDirectoryContent(w http.ResponseWriter, r *http.Request) {
	directoryPath := r.PathValue("path")
	directoryPath = strings.ReplaceAll(directoryPath, "%20", " ")

	directoryContent, err := filesystem.NewDirectoryContent(filepath.Join(s.config.MarkdownDirectory, directoryPath))
	if err != nil {
		logging.Error(err.Error())
		s.respondWithError(w, http.StatusInternalServerError, internal.OnyxError{
			ErrorMessage: fmt.Sprintf("unable to read directory content: %s", directoryPath),
		})
		return
	}

	err = json.NewEncoder(w).Encode(directoryContent)
	if err != nil {
		logging.Error(err.Error())
		s.respondWithError(w, http.StatusInternalServerError, internal.OnyxError{
			ErrorMessage: fmt.Sprintf("unable to encode directory content to JSON: %s", directoryPath),
		})
		return
	}
}
