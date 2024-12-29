package onyx

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/pazifical/onyx/internal/database"
	"github.com/pazifical/onyx/logging"
)

type Server struct {
	config         Config
	mux            *http.ServeMux
	noteRepository *database.NoteRepository
	frontendFS     embed.FS
}

func NewServer(config Config, frontendFS embed.FS) *Server {
	noteRepo := database.NewNoteRepository(config.MarkdownDirectory)

	server := Server{
		config:         config,
		mux:            http.NewServeMux(),
		noteRepository: &noteRepo,
		frontendFS:     frontendFS,
	}

	server.mux.HandleFunc("GET /", server.ServeIndex)
	server.mux.HandleFunc("GET /favicon.ico", server.ServeFavIcon)
	server.mux.HandleFunc("GET /assets/{path...}", server.ServeAsset)
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
