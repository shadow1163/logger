package logger

import "testing"

func TestDebug(t *testing.T) {
	logger := NewLogger()
	logger.Debug("123", 123, 12, "123")
}

func TestInfo(t *testing.T) {
	logger := NewLogger()
	logger.Info("123", 123, 12, "123")
}

func TestWarning(t *testing.T) {
	logger := NewLogger()
	logger.Warning("123", 123, 12, "123")
}

func TestError(t *testing.T) {
	logger := NewLogger()
	logger.Error("123", 123, 12, "123")
}

func TestPrintln(t *testing.T) {
	logger := NewLogger()
	logger.Println("testing")
}
