package logger

import (
	"log"
)

// Log ...
type Log struct{}

// Error ...
func (l *Log) Error(args ...interface{}) {
	log.Println(args...)
}

// Errorf ...
func (l *Log) Errorf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

// Fatalf ...
func (l *Log) Fatalf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

// Fatal ...
func (l *Log) Fatal(args ...interface{}) {
	log.Println(args...)
}

// Infof ...
func (l *Log) Infof(format string, args ...interface{}) {
	log.Printf(format, args...)
}

// Info ...
func (l *Log) Info(args ...interface{}) {
	log.Println(args...)
}

// Warnf ...
func (l *Log) Warnf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

// Warn ...
func (l *Log) Warn(args ...interface{}) {
	log.Println(args...)
}

// Debugf ...
func (l *Log) Debugf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

// Debug ...
func (l *Log) Debug(args ...interface{}) {
	log.Println(args...)
}

// Panicf ...
func (l *Log) Panicf(format string, args ...interface{}) {
	log.Printf(format, args...)
	panic("")
}

// Panic ...
func (l *Log) Panic(args ...interface{}) {
	log.Println(args...)
	panic("")
}

// With ...
func (l *Log) With(val ...interface{}) Logger {
	return &Log{}
}
