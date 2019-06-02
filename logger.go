package logger

import (
	"log"
	"os"
)

var Log *logger = NewLogger()

type logger struct {
	log *log.Logger
}

func NewLogger() *logger {
	logger := logger{}
	logger.log = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	return &logger
}

func (logger *logger) Debug(message ...interface{}) {
	logger.log.SetPrefix("[DEBUG] ")
	logger.log.Println(message...)
}

func (logger *logger) Info(message ...interface{}) {
	logger.log.SetPrefix("[INFO] ")
	logger.log.Println(message...)
}

func (logger *logger) Warning(message ...interface{}) {
	logger.log.SetPrefix("[Warning] ")
	logger.log.Println(message...)
}

func (logger *logger) Error(message ...interface{}) {
	logger.log.SetPrefix("[ERROR] ")
	logger.log.Println(message...)
}
