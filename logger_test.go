package logger

import "testing"

func TestDebug(t *testing.T) {
	logger := NewLogger()
	logger.SetLevel(DEBUG)
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

func TestLevel(t *testing.T) {
	logger := NewLogger()
	logger.SetLevel(10)
	logger.Debug("Debug")
	logger.Info("Info")
	logger.Warning("Warning")
	logger.Error("Error")
}

func TestOutput(t *testing.T) {
	logger := NewLogger()
	logger.Stdf("output")
	logger.Stdf("and")
}

func TestPrintF(t *testing.T) {
	logger := NewLogger()
	logger.SetLevel(DEBUG)
	logger.NOCOLOR = true
	aaa := "aaaa"
	logger.Errorf("%s", aaa)
	logger.Warningf("%s", aaa)
	logger.Infof("%s", aaa)
	logger.Debugf("%s", aaa)
}
