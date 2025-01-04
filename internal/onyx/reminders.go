package onyx

import (
	"encoding/json"
	"net/http"

	"github.com/pazifical/onyx/internal"
	"github.com/pazifical/onyx/logging"
)

func (s *Server) GetAllReminders(w http.ResponseWriter, r *http.Request) {
	reminders := s.monitoringService.GetAllReminders()
	err := json.NewEncoder(w).Encode(reminders)
	if err != nil {
		logging.Error(err.Error())
		s.respondWithError(w, http.StatusInternalServerError, internal.OnyxError{
			ErrorMessage: "unable to encode reminders to JSON",
		})
		return
	}
}
