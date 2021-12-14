package autoeq

import (
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

type MyLogger struct {
	*logrus.Logger
}

type ErrorEvent struct {
	id  uint
	msg string
}

var defaultFields = logrus.Fields{
	"appname":    "Go-AutoEQ",
	"go-version": runtime.Version(),
}

var (
	successEvent           = &ErrorEvent{0, "success: %s"}
	fileNotOpenedEvent     = &ErrorEvent{1, "couldn't open the requested file: %s"}
	configNotExportedEvent = &ErrorEvent{2, "couldn't export the config profile file: %s"}
	invalidArgRangeEvent   = &ErrorEvent{3, "argument %s is out of expected bounds"}
)

func newWrapper(logger *logrus.Logger) *MyLogger {
	return &MyLogger{logger}
}

func newBaseLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)

	return logger
}

func NewLogger() *MyLogger {
	baseLogger := newBaseLogger()
	wrapper := newWrapper(baseLogger)

	return wrapper
}

func (l *MyLogger) SuccessLog(context string) {
	l.WithFields(defaultFields).Infof(successEvent.msg, context)
}

func (l *MyLogger) FileNotOpenedLog(path string) {
	l.WithFields(defaultFields).Errorf(fileNotOpenedEvent.msg, path)
}

func (l *MyLogger) ConfigNotExportedLog(path string) {
	l.WithFields(defaultFields).Errorf(configNotExportedEvent.msg, path)
}

func (l *MyLogger) InvalidArgRangeLog(arg string) {
	l.WithFields(defaultFields).Warnf(invalidArgRangeEvent.msg, arg)
}
