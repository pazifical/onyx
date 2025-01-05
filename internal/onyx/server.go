package onyx

import (
	"embed"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pazifical/onyx/internal"
	"github.com/pazifical/onyx/internal/database"
	"github.com/pazifical/onyx/internal/filesystem"
	"github.com/pazifical/onyx/internal/matrix"
	"github.com/pazifical/onyx/internal/reminder"
	"github.com/pazifical/onyx/logging"
)

type Server struct {
	config            Config
	mux               *http.ServeMux
	noteRepository    *database.NoteRepository
	frontendFS        embed.FS
	monitoringService *reminder.MonitoringService
}

func NewServer(config Config, frontendFS embed.FS) *Server {
	noteRepo := database.NewNoteRepository(config.MarkdownDirectory)

	monitoringService := reminder.NewMonitoringService(&noteRepo)

	server := Server{
		config:            config,
		mux:               http.NewServeMux(),
		noteRepository:    &noteRepo,
		frontendFS:        frontendFS,
		monitoringService: &monitoringService,
	}

	server.mux.HandleFunc("GET /", server.ServeIndex)
	server.mux.HandleFunc("GET /favicon.ico", server.ServeFavIcon)
	server.mux.HandleFunc("GET /assets/{path...}", server.ServeAsset)
	server.mux.HandleFunc("GET /api/notes", server.GetAllNotes)
	server.mux.HandleFunc("GET /api/notes/{path...}", server.GetNoteByFilePath)
	server.mux.HandleFunc("PUT /api/notes/{path...}", server.SaveNote)
	server.mux.HandleFunc("POST /api/notes/{path...}", server.SaveNote)
	server.mux.HandleFunc("GET /api/directory/{path...}", server.GetDirectoryContent)
	server.mux.HandleFunc("POST /api/directory/{path...}", server.CreateDirectory)
	server.mux.HandleFunc("GET /api/reminders", server.GetAllReminders)
	server.mux.HandleFunc("GET /api/directory_tree", server.GetDirectoryTree)

	return &server
}

func (s *Server) AddMatrixService(matrixService *matrix.Service) {
	s.monitoringService.InitializeMatrixService(matrixService)
}

func (s *Server) Start() error {
	go s.monitoringService.Start()

	logging.Info(fmt.Sprintf("starting Onyx server on port %d", s.config.Port))

	err := http.ListenAndServe(fmt.Sprintf(":%d", s.config.Port), s.mux)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) respondWithError(w http.ResponseWriter, httpStatusCode int, onyxError internal.OnyxError) {
	w.WriteHeader(httpStatusCode)

	err := json.NewEncoder(w).Encode(onyxError)
	if err != nil {
		logging.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) GetDirectoryTree(w http.ResponseWriter, r *http.Request) {
	dirTree, err := filesystem.CreateDirectoryTree(s.config.MarkdownDirectory)
	if err != nil {
		s.respondWithError(w, http.StatusInternalServerError, internal.OnyxError{ErrorMessage: "unable to create a directory tree"})
		return
	}

	err = json.NewEncoder(w).Encode(dirTree)
	if err != nil {
		s.respondWithError(w, http.StatusInternalServerError, internal.OnyxError{ErrorMessage: "unable to encode directory tree to JSON"})
		return
	}
}
