package onyx

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pazifical/onyx/logging"
)

func (s *Server) ServeIndex(w http.ResponseWriter, r *http.Request) {
	data, err := s.frontendFS.ReadFile("frontend/dist/index.html")
	if err != nil {
		logging.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		logging.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) ServeFavIcon(w http.ResponseWriter, r *http.Request) {
	data, err := s.frontendFS.ReadFile("frontend/dist/favicon.ico")
	if err != nil {
		logging.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		logging.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) ServeAsset(w http.ResponseWriter, r *http.Request) {
	filePath := r.PathValue("path")
	assetPath := fmt.Sprintf("frontend/dist/assets/%s", filePath)

	data, err := s.frontendFS.ReadFile(assetPath)
	if err != nil {
		logging.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
