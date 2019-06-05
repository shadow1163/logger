package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

// Log instance
var Log = NewLogger()

// Logger struct
type Logger struct {
	*log.Logger
}

//NewLogger Create log instance
func NewLogger() *Logger {
	log := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	logger := Logger{log}
	return &logger
}

// Debug print debug
func (logger *Logger) Debug(message ...interface{}) {
	logger.setPrefix("DEBUG")
	logger.Println(message...)
}

// Info print info
func (logger *Logger) Info(message ...interface{}) {
	logger.setPrefix("INFO")
	logger.Println(message...)
}

// Warning print warning
func (logger *Logger) Warning(message ...interface{}) {
	logger.setPrefix("Warning")
	logger.Println(message...)
}

// Error print error
func (logger *Logger) Error(message ...interface{}) {
	logger.setPrefix("ERROR")
	logger.Println(message...)
}

func (logger *Logger) caller(skip int) (pc uintptr, file string, line int, ok bool) {
	return runtime.Caller(skip)
}

func (logger *Logger) setPrefix(level string) {
	pc, file, line, ok := logger.caller(3)
	if !ok {
		logger.Println("Error to get func name and file name!")
		return
	}
	logger.SetPrefix(fmt.Sprintf("[%s] %s %s %d ", level, file, runtime.FuncForPC(pc).Name(), line))
}
