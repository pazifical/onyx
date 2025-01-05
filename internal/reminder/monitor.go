package reminder

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/pazifical/onyx/internal/database"
	"github.com/pazifical/onyx/internal/matrix"
	"github.com/pazifical/onyx/internal/types"
	"github.com/pazifical/onyx/logging"
)

var interval = time.Hour
var onyxDirectiveRegex *regexp.Regexp

func init() {
	r, err := regexp.Compile(`\[\[.*:.*\]\]`)
	if err != nil {
		log.Fatal(err)
	}

	onyxDirectiveRegex = r
}

type MonitoringService struct {
	repository    *database.NoteRepository
	reminders     []Reminder
	matrixService *matrix.Service
}

func NewMonitoringService(repository *database.NoteRepository) MonitoringService {
	return MonitoringService{
		repository: repository,
	}
}

func (ms *MonitoringService) InitializeMatrixService(matrixService *matrix.Service) {
	ms.matrixService = matrixService
}

func (ms *MonitoringService) GetAllReminders() []Reminder {
	return ms.reminders
}

func (ms *MonitoringService) Start() {
	logging.Info("staring Onyx reminder monitoring service")
	for {
		logging.Info("searching markdown files for reminders")
		ms.reminders = make([]Reminder, 0)
		err := ms.searchForReminders()
		if err != nil {
			logging.Error(err.Error())
		}
		time.Sleep(interval)
	}
}

func (ms *MonitoringService) searchForReminders() error {
	notes, err := ms.repository.FetchAll()
	if err != nil {
		return err
	}

	for _, note := range notes {
		reminders := extractRemindersFromNote(note)
		if len(reminders) > 0 {
			ms.reminders = append(ms.reminders, reminders...)
		}
	}

	if len(ms.reminders) == 0 {
		return nil
	}

	if ms.matrixService != nil {
		err = ms.matrixService.Authenticate()
		if err != nil {
			logging.Error(err.Error())
		}

		err = ms.sendMatrixMessages(ms.reminders)
		if err != nil {
			logging.Error(err.Error())
		}
	}

	return nil
}

func (ms *MonitoringService) sendMatrixMessages(reminders []Reminder) error {
	err := ms.matrixService.Authenticate()
	if err != nil {
		return err
	}

	var builder strings.Builder
	for _, reminder := range reminders {
		builder.WriteString(fmt.Sprintf("%s\n%s\n(%s)\n\n\n", reminder.Date, reminder.Content, reminder.Source))
	}

	err = ms.matrixService.SendMessage(builder.String())
	if err != nil {
		return err
	}

	return nil
}

func extractRemindersFromNote(note types.Note) []Reminder {
	reminders := make([]Reminder, 0)

	for _, line := range strings.Split(note.Text, "\n") {
		reminder, ok := extractReminderFromLine(line, note.Path)
		if !ok {
			continue
		}
		reminders = append(reminders, reminder)

		logging.Info(fmt.Sprintf("found reminder (%s): %s", reminder.Type, reminder.Content))
	}
	return reminders
}

func extractOnyxExpression(text string) (OnyxExpression, bool) {
	match := onyxDirectiveRegex.FindString(text)
	if match == "" {
		return OnyxExpression{}, false
	}

	stripped := strings.Replace(match, "[[", "", 1)
	stripped = strings.Replace(stripped, "]]", "", 1)

	parts := strings.SplitN(stripped, ":", 2)
	if len(parts) != 2 {
		return OnyxExpression{}, false
	}

	return OnyxExpression{
		Type:    parts[0],
		Content: parts[1],
	}, true
}

func extractReminderFromLine(line string, source string) (Reminder, bool) {
	prefix := "- [ ] "
	if !strings.HasPrefix(line, prefix) {
		return Reminder{}, false
	}

	stripped := strings.Replace(line, prefix, "", 1)

	onyxExpr, ok := extractOnyxExpression(stripped)
	if !ok {
		return Reminder{}, false
	}

	if onyxExpr.Type == deadlineType {
		date, err := time.Parse("2006-01-02", onyxExpr.Content)
		if err != nil {
			logging.Error(err.Error())
			return Reminder{}, false
		}

		if time.Now().Before(date) {
			return Reminder{}, false
		}

		return Reminder{
			Date:    date,
			Content: strings.TrimSpace(stripped),
			Type:    onyxExpr.Type,
			Source:  source,
		}, true
	} else if onyxExpr.Type == birthdayType {
		date, err := time.Parse("2006-01-02", onyxExpr.Content)
		if err != nil {
			logging.Error(err.Error())
			return Reminder{}, false
		}

		now := time.Now()
		if now.Day() != date.Day() || now.Month() != date.Month() {
			return Reminder{}, false
		}

		return Reminder{
			Date:    date,
			Content: strings.TrimSpace(stripped),
			Type:    onyxExpr.Type,
			Source:  source,
		}, true
	} else if onyxExpr.Type == dateType {
		date, err := time.Parse("2006-01-02", onyxExpr.Content)
		if err != nil {
			logging.Error(err.Error())
			return Reminder{}, false
		}

		now := time.Now()
		if now.Day() != date.Day() || now.Month() != date.Month() || now.Year() != date.Year() {
			return Reminder{}, false
		}

		return Reminder{
			Date:    date,
			Content: strings.TrimSpace(stripped),
			Type:    onyxExpr.Type,
			Source:  source,
		}, true
	} else if onyxExpr.Type == untilType {
		date, err := time.Parse("2006-01-02", onyxExpr.Content)
		if err != nil {
			logging.Error(err.Error())
			return Reminder{}, false
		}

		if time.Now().After(date) {
			return Reminder{}, false
		}

		return Reminder{
			Date:    date,
			Content: strings.TrimSpace(stripped),
			Type:    onyxExpr.Type,
			Source:  source,
		}, true
	} else {
		logging.Warning(fmt.Sprintf("onyx expression '%s' not implemented yet", onyxExpr.Type))
	}

	return Reminder{}, false
}
