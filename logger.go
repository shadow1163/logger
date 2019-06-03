package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

// Log instance
var Log = NewLogger()

type logger struct {
	log *log.Logger
}

//NewLogger Create log instance
func NewLogger() *logger {
	logger := logger{}
	logger.log = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	return &logger
}

func (logger *logger) Debug(message ...interface{}) {
	logger.setPrefix("DEBUG")
	logger.log.Println(message...)
}

func (logger *logger) Info(message ...interface{}) {
	logger.setPrefix("INFO")
	logger.log.Println(message...)
}

func (logger *logger) Warning(message ...interface{}) {
	logger.setPrefix("Warning")
	logger.log.Println(message...)
}

func (logger *logger) Error(message ...interface{}) {
	logger.setPrefix("ERROR")
	logger.log.Println(message...)
}

func (logger *logger) caller(skip int) (pc uintptr, file string, line int, ok bool) {
	return runtime.Caller(skip)
}

func (logger *logger) setPrefix(level string) {
	pc, file, line, ok := logger.caller(3)
	if !ok {
		logger.log.Println("Error to get func name and file name!")
		return
	}
	logger.log.SetPrefix(fmt.Sprintf("[%s] %s %s %d ", level, file, runtime.FuncForPC(pc).Name(), line))
}
