package onyx

import (
	"encoding/json"
	"net/http"

	"github.com/pazifical/onyx/internal"
	"github.com/pazifical/onyx/internal/types"
	"github.com/pazifical/onyx/logging"
)

func (s *Server) GetAllReminders(w http.ResponseWriter, r *http.Request) {
	reminders := s.reminderService.GetAllReminders()
	err := json.NewEncoder(w).Encode(reminders)
	if err != nil {
		logging.Error(err.Error())
		s.respondWithError(w, http.StatusInternalServerError, internal.OnyxError{
			ErrorMessage: "unable to encode reminders to JSON",
		})
		return
	}
}

func (s *Server) AddReminder(w http.ResponseWriter, r *http.Request) {
	var reminder types.Reminder
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&reminder)
	if err != nil {
		logging.Error(err.Error())
		s.respondWithError(w, http.StatusNotAcceptable, internal.OnyxError{
			ErrorMessage: "unable to decode reminder from JSON",
		})
		return
	}

	err = s.reminderService.AddReminder(reminder)
	if err != nil {
		logging.Error(err.Error())
		s.respondWithError(w, http.StatusNotAcceptable, internal.OnyxError{
			ErrorMessage: "unable to add reminder",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
}
