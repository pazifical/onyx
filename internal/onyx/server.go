package onyx

import (
	"fmt"
	"net/http"

	"github.com/pazifical/onyx/internal/database"
	"github.com/pazifical/onyx/logging"
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
	server.mux.HandleFunc("GET /api/directory/{path...}", server.GetDirectoryContent)

	return &server
}

func (s *Server) Start() error {
	logging.Info(fmt.Sprintf("starting Onyx server on port %d", s.config.Port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", s.config.Port), s.mux)
	if err != nil {
		return err
	}

	return nil
}
