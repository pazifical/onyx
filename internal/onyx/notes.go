package onyx

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/pazifical/onyx/internal/filesystem"
	"github.com/pazifical/onyx/internal/types"
	"github.com/pazifical/onyx/logging"
)

func (s *Server) GetAllNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := s.noteRepository.FetchAll()
	if err != nil {
		logging.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(notes)
	if err != nil {
		logging.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) GetNoteByID(w http.ResponseWriter, r *http.Request) {
	noteID := r.PathValue("id")

	note, err := s.noteRepository.FetchOne(noteID)
	if err != nil {
		logging.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(note)
	if err != nil {
		logging.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) SaveNote(w http.ResponseWriter, r *http.Request) {
	var note types.Note

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		logging.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	note, err = s.noteRepository.Update(note)
	if err != nil {
		logging.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(note)
	if err != nil {
		logging.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func (s *Server) GetNoteByFilePath(w http.ResponseWriter, r *http.Request) {
	filePath := r.PathValue("path")
	filePath = strings.ReplaceAll(filePath, "%20", " ")

	note, err := s.noteRepository.FetchOne(filePath)
	if err != nil {
		logging.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(note)
	if err != nil {
		logging.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) GetDirectoryContent(w http.ResponseWriter, r *http.Request) {
	directoryPath := r.PathValue("path")
	directoryPath = strings.ReplaceAll(directoryPath, "%20", " ")

	directoryContent, err := filesystem.NewDirectoryContent(filepath.Join(s.config.MarkdownDirectory, directoryPath))
	if err != nil {
		logging.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(directoryContent)
	if err != nil {
		logging.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
