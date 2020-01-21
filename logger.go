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
	Level int
	*log.Logger
}

const (
	NONE = iota //0 NONE
	ERROR
	WARNING
	INFO
	DEBUG
)

//NewLogger Create log instance
func NewLogger() *Logger {
	log := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	logger := Logger{INFO, log}
	logger.setPrefix("INFO")
	return &logger
}

// Debug print debug
func (logger *Logger) Debug(message ...interface{}) {
	if logger.Level >= DEBUG {
		logger.setPrefix("DEBUG")
		logger.Println(message...)
	}
}

// Debugf print debug
func (logger *Logger) Debugf(format string, message ...interface{}) {
	if logger.Level >= DEBUG {
		logger.setPrefix("DEBUG")
		logger.Println(fmt.Sprintf(format, message...))
	}
}

// Info print info
func (logger *Logger) Info(message ...interface{}) {
	if logger.Level >= INFO {
		logger.setPrefix("INFO")
		logger.Println(message...)
	}
}

// Infof print info
func (logger *Logger) Infof(format string, message ...interface{}) {
	if logger.Level >= INFO {
		logger.setPrefix("INFO")
		logger.Printf(format, message...)
	}
}

// Warning print warning
func (logger *Logger) Warning(message ...interface{}) {
	if logger.Level >= WARNING {
		logger.setPrefix("Warning")
		logger.Println(message...)
	}
}

// Warningf print warning
func (logger *Logger) Warningf(format string, message ...interface{}) {
	if logger.Level >= WARNING {
		logger.setPrefix("Warning")
		logger.Printf(format, message...)
	}
}

// Error print error
func (logger *Logger) Error(message ...interface{}) {
	if logger.Level >= ERROR {
		logger.setPrefix("ERROR")
		logger.Println(message...)
	}
}

// Errorf print error
func (logger *Logger) Errorf(format string, message ...interface{}) {
	if logger.Level >= ERROR {
		logger.setPrefix("ERROR")
		logger.Printf(format, message...)
	}
}

// SetLevel set log level
func (logger *Logger) SetLevel(level int) {
	if level >= 5 || level < 0 {
		logger.Error("Invalid Level, must one of NONE, ERROR, WARNING, INFO, DEBUG!")
		return
	}
	logger.Level = level
}

// Std no level to print
func (logger *Logger) Std(message ...interface{}) {
	fmt.Print(message...)
}

// Stdf no level to print
func (logger *Logger) Stdf(format string, message ...interface{}) {
	fmt.Printf(format, message...)
}

func (logger *Logger) caller(skip int) (pc uintptr, file string, line int, ok bool) {
	return runtime.Caller(skip)
}

func (logger *Logger) setPrefix(level string) {
	_, file, line, ok := logger.caller(3)
	if !ok {
		logger.Println("Error to get func name and file name!")
		return
	}
	// logger.SetPrefix(fmt.Sprintf("[%s] %s %s %d ", level, file, runtime.FuncForPC(pc).Name(), line))
	logger.SetPrefix(fmt.Sprintf("[%s] %s %d ", level, file, line))
}
