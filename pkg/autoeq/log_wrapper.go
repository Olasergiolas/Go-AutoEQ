package autoeq

import (
	"errors"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

type MyLogger struct {
	*logrus.Logger
}

type ErrorType uint

const (
	Success = iota
	ParametricDataNotFound
	Undefined
)

var (
	eventsMsg = map[ErrorType]error{
		Success:                errors.New("success! %d"),
		ParametricDataNotFound: errors.New("error %d, couldn't open the file containing the param. data"),
		Undefined:              errors.New("unkown event code %d"),
	}

	defaultFields = logrus.Fields{
		"appname":    "Go-AutoEQ",
		"go-version": runtime.Version(),
	}
)

func NewWrapper(logger *logrus.Logger) *MyLogger {
	return &MyLogger{logger}
}

func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)

	return logger
}

func (logger MyLogger) Log(errCode uint) {
	msg, ok := eventsMsg[ErrorType(errCode)]
	code := errCode

	if !ok {
		msg = eventsMsg[Undefined]
		code = Undefined
	}
	logger.WithFields(defaultFields).Infof(msg.Error(), code)
}
