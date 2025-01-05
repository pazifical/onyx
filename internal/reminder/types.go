package reminder

import "time"

const birthdayType = "birthday"
const deadlineType = "deadline"
const dateType = "date"
const untilType = "until"

type Reminder struct {
	Date    time.Time `json:"date"`
	Content string    `json:"todo"`
	Type    string    `json:"string"`
	Source  string    `json:"source"`
}

type OnyxExpression struct {
	Type    string
	Content string
}
