package onyx

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pazifical/onyx/internal"
	"github.com/pazifical/onyx/logging"
)

func (s *Server) ServeIndex(w http.ResponseWriter, r *http.Request) {
	data, err := s.frontendFS.ReadFile("frontend/dist/index.html")
	if err != nil {
		logging.Error(err.Error())
		s.respondWithError(w, http.StatusInternalServerError, internal.OnyxError{
			ErrorMessage: "unable to read frontend index.html",
		})
		return
	}

	_, err = w.Write(data)
	if err != nil {
		logging.Error(err.Error())
		s.respondWithError(w, http.StatusInternalServerError, internal.OnyxError{
			ErrorMessage: "unable to serve frontend",
		})
		return
	}
}

func (s *Server) ServeFavIcon(w http.ResponseWriter, r *http.Request) {
	data, err := s.frontendFS.ReadFile("frontend/dist/favicon.ico")
	if err != nil {
		logging.Error(err.Error())
		s.respondWithError(w, http.StatusNotFound, internal.OnyxError{
			ErrorMessage: "unable to read frontend favicon",
		})
		return
	}

	_, err = w.Write(data)
	if err != nil {
		logging.Error(err.Error())
		s.respondWithError(w, http.StatusInternalServerError, internal.OnyxError{
			ErrorMessage: "unable to serve frontend",
		})
		return
	}
}

func (s *Server) ServeAsset(w http.ResponseWriter, r *http.Request) {
	filePath := r.PathValue("path")
	assetPath := fmt.Sprintf("frontend/dist/assets/%s", filePath)

	data, err := s.frontendFS.ReadFile(assetPath)
	if err != nil {
		logging.Error(err.Error())
		s.respondWithError(w, http.StatusNotFound, internal.OnyxError{
			ErrorMessage: fmt.Sprintf("unable to read frontend asset '%s'", filePath),
		})
		return
	}

	if strings.HasSuffix(assetPath, ".js") {
		w.Header().Set("Content-Type", "application/javascript")
	} else if strings.HasSuffix(assetPath, ".css") {
		w.Header().Set("Content-Type", "text/css")
	}

	_, err = w.Write(data)
	if err != nil {
		logging.Error(err.Error())
		s.respondWithError(w, http.StatusNotFound, internal.OnyxError{
			ErrorMessage: fmt.Sprintf("unable to serve frontend asset '%s'", filePath),
		})
		return
	}
}
