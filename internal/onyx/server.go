package onyx

import (
	"fmt"
	"net/http"

	"github.com/pazifical/onyx/internal/database"
)

type Server struct {
	config         Config
	mux            *http.ServeMux
	noteRepository *database.NoteRepository
}

func NewServer(config Config) *Server {
	noteRepo := database.NewNoteRepository(config.MarkdownDirectory)

	server := Server{
		config:         config,
		mux:            http.NewServeMux(),
		noteRepository: &noteRepo,
	}

	server.mux.HandleFunc("GET /api/notes", server.GetAllNotes)
	server.mux.HandleFunc("GET /api/notes/{path...}", server.GetNoteByFilePath)

	return &server
}

func (s *Server) Start() error {
	err := http.ListenAndServe(fmt.Sprintf(":%d", s.config.Port), s.mux)
	if err != nil {
		return err
	}

	return nil
}
