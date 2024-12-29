package logging

import "log"

func Info(message string) {
	log.Printf("| INFO:    %s", message)
}
func Warning(message string) {
	log.Printf("| WARNING: %s", message)
}
func Error(message string) {
	log.Printf("| ERROR:   %s", message)
}
func Debug(message string) {
	log.Printf("| DEBUG:   %s", message)
}
