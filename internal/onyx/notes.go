package onyx

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	fmt.Println(filePath)

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
