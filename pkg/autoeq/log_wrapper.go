package autoeq

import (
	"errors"
	"os"

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
		Success:                errors.New("Success! %d"),
		ParametricDataNotFound: errors.New("Error %d, couldn't open the file containing the param. data"),
		Undefined:              errors.New("Unkown event code %d"),
	}

	defaultFields = logrus.Fields{
		"appname": "Go-AutoEQ",
	}
)

func NewLogger() MyLogger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)

	return MyLogger{logger}
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
