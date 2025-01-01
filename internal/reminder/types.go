package reminder

import "time"

type Reminder struct {
	Date   time.Time `json:"date"`
	ToDo   string    `json:"todo"`
	Source string    `json:"source"`
}
