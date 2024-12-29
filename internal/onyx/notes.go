package onyx

import (
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/pazifical/onyx/internal/filesystem"
)

func (s *Server) GetAllNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := s.noteRepository.FetchAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(notes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) GetNoteByID(w http.ResponseWriter, r *http.Request) {
	noteID := r.PathValue("id")

	note, err := s.noteRepository.FetchOne(noteID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) GetNoteByFilePath(w http.ResponseWriter, r *http.Request) {
	filePath := r.PathValue("path")

	note, err := s.noteRepository.FetchOne(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) GetDirectoryContent(w http.ResponseWriter, r *http.Request) {
	directoryPath := r.PathValue("path")

	directoryContent, err := filesystem.NewDirectoryContent(filepath.Join(s.config.MarkdownDirectory, directoryPath))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(directoryContent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
