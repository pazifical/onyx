package reminder

import (
	"testing"
	"time"
)

func TestExtractReminder1(t *testing.T) {
	line := "- [ ] Learn something (#until:2024-12-24)"

	wantSource := "some source"

	reminder, ok := extractReminderFromLine(line, wantSource)
	if !ok {
		t.Errorf("No reminder could be extracted from line '%s'", line)
	}

	wantToDo := "Learn something"
	if reminder.ToDo != wantToDo {
		t.Errorf("%s != %s", reminder.ToDo, wantToDo)
	}

	wantDate := time.Date(2024, time.December, 24, 0, 0, 0, 0, time.UTC)
	if reminder.Date != wantDate {
		t.Errorf("%v != %v", reminder.Date, wantDate)
	}

	if wantSource != reminder.Source {
		t.Errorf("%s != %s", reminder.Source, wantSource)
	}
}

func TestExtractReminder2(t *testing.T) {
	line := "- [ ] Learn something (#until:2024-12-24) and something else"

	wantSource := "some source"

	reminder, ok := extractReminderFromLine(line, wantSource)
	if !ok {
		t.Errorf("No reminder could be extracted from line '%s'", line)
	}

	// TODO: Should this really be the behaviour?
	wantToDo := "Learn something  and something else"
	if reminder.ToDo != wantToDo {
		t.Errorf("%s != %s", reminder.ToDo, wantToDo)
	}

	wantDate := time.Date(2024, time.December, 24, 0, 0, 0, 0, time.UTC)
	if reminder.Date != wantDate {
		t.Errorf("%v != %v", reminder.Date, wantDate)
	}

	if wantSource != reminder.Source {
		t.Errorf("%s != %s", reminder.Source, wantSource)
	}
}

func TestExtractReminder3(t *testing.T) {
	line := "- [ ] Learn something (#until:10024-12-24)"

	wantSource := "some source"

	_, ok := extractReminderFromLine(line, wantSource)
	if ok {
		t.Error("A reminder should not be created because the event is not due yet")
	}

}
