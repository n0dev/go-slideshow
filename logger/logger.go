package logger

import (
	"log"
	"os"
)

const (
	logName  = "goslideshow.log"
	logFlags = log.Ldate | log.Ltime
)

var (
	trace   *log.Logger
	info    *log.Logger
	warning *log.Logger
	error   *log.Logger
)

// Create the logger file before everything
func init() {
	fileFlags := os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	file, err := os.OpenFile(logName, fileFlags, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", os.Stdout, ":", err)
	}

	trace = log.New(file, "TRACE: ", logFlags)
	info = log.New(file, "INFO: ", logFlags)
	warning = log.New(file, "WARNING: ", logFlags)
	error = log.New(file, "ERROR: ", logFlags)

	Trace("Starting GoSlideshow")
}

// Trace prints the message with newline in the logfile
func Trace(message string) {
	trace.Println(message)
}

// Info prints the message with newline in the logfile
func Info(message string) {
	info.Println(message)
}

// Warning prints the message with newline in the logfile
func Warning(message string) {
	warning.Println(message)
}

// Error prints the message with newline in the logfile
func Error(message string) {
	error.Println(message)
}
