package onyx

import (
	"encoding/json"
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
