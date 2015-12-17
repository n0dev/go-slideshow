package logger

import (
	"os"
	"time"
)

// Logger is the logger file
type Logger struct {
	fd *os.File
}

// New is to initialize the logger
func New(name string) (*Logger, error) {
	var log Logger
	var err error
	log.fd, err = os.Create(name)
	if err != nil {
		panic("Cannot create logfile")
	}
	return &log, nil
}

func (l *Logger) log(level, message string) {
	t := time.Now()
	l.fd.WriteString(t.Format("2006-01-02 15:04:05") + " - " + level + " - " + message + "\n")
}

// Panic is just like func l.Critical except that it is followed by a call to panic
func (l *Logger) Panic(message string) {
	l.log("CRITICAL", message)
	panic(message)
}

// Critical logs a message at a Critical Level
func (l *Logger) Critical(message string) {
	l.log("CRITICAL", message)
}

// Error logs a message at Error level
func (l *Logger) Error(message string) {
	l.log("ERROR", message)
}

// Warning logs a message at Warning level
func (l *Logger) Warning(message string) {
	l.log("WARNING", message)
}

// Info logs a message at Info level
func (l *Logger) Info(message string) {
	l.log("INFO", message)
}

// Debug logs a message at Debug level
func (l *Logger) Debug(message string) {
	l.log("DEBUG", message)
}
