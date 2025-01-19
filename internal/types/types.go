package types

import "time"

type Note struct {
	Path string `json:"path"`
	Text string `json:"text"`
}

type Reminder struct {
	Date    time.Time `json:"date"`
	Content string    `json:"todo"`
	Type    string    `json:"string"`
	Source  string    `json:"source"`
}
