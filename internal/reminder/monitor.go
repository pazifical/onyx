package reminder

import (
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/pazifical/onyx/internal/database"
	"github.com/pazifical/onyx/internal/types"
	"github.com/pazifical/onyx/logging"
)

var interval = time.Hour
var untilRegex *regexp.Regexp

func init() {
	r, err := regexp.Compile(`\(#until:\d\d\d\d-\d{1,2}-\d{1,2}\)`)
	if err != nil {
		log.Fatal(err)
	}

	untilRegex = r
}

type MonitoringService struct {
	repository *database.NoteRepository
	reminders  []Reminder
}

func NewMonitoringService(repository *database.NoteRepository) MonitoringService {
	return MonitoringService{
		repository: repository,
	}
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
	}
	return reminders
}

func extractReminderFromLine(line string, source string) (Reminder, bool) {
	prefix := "- [ ] "
	if !strings.HasPrefix(line, prefix) {
		return Reminder{}, false
	}

	stripped := strings.Replace(line, prefix, "", 1)

	match := untilRegex.FindString(line)
	if match == "" {
		return Reminder{}, false
	}

	todo := strings.Replace(stripped, match, "", 1)

	dateString := strings.Replace(match, "(#until:", "", 1)
	dateString = strings.Replace(dateString, ")", "", 1)

	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		logging.Error(err.Error())
		return Reminder{}, false
	}

	if time.Now().Before(date) {
		return Reminder{}, false
	}

	return Reminder{
		Date:   date,
		ToDo:   strings.TrimSpace(todo),
		Source: source,
	}, true
}
