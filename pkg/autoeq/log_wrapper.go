package autoeq

import (
	"os"
	"runtime"

	"github.com/rs/zerolog"
)

type MyLogger struct {
	*zerolog.Logger
}

type ErrorEvent struct {
	id  uint
	msg string
}

var defaultFields = map[string]interface{}{
	"appname":    "Go-AutoEQ",
	"go-version": runtime.Version(),
}

var (
	successEvent           = &ErrorEvent{0, "success: %s"}
	fileNotOpenedEvent     = &ErrorEvent{1, "couldn't open the requested file: %s"}
	configNotExportedEvent = &ErrorEvent{2, "couldn't export the config profile file: %s"}
	invalidArgRangeEvent   = &ErrorEvent{3, "argument %s is out of expected bounds"}
	missingEnvVarEvent     = &ErrorEvent{4, "env variable %s does not exist"}
	configInfoEvent        = &ErrorEvent{5, "var %s: value %s"}
)

func newWrapper(logger *zerolog.Logger) *MyLogger {
	return &MyLogger{logger}
}

func newBaseLogger() *zerolog.Logger {
	var logger zerolog.Logger
	f, err := prepareLogFile()

	if err == nil {
		logger = zerolog.New(f)
	} else {
		logger = zerolog.New(os.Stdout)
	}
	return &logger
}

func NewLogger() *MyLogger {
	baseLogger := newBaseLogger()
	wrapper := newWrapper(baseLogger)

	return wrapper
}

func prepareLogFile() (*os.File, error) {
	path := GetLogPath()
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		return nil, err
	}

	return f, nil
}

func (l *MyLogger) SuccessLog(context string) {
	l.Info().Timestamp().Fields(defaultFields).Msgf(successEvent.msg, context)
}

func (l *MyLogger) FileNotOpenedLog(path string) {
	l.Error().Timestamp().Fields(defaultFields).Msgf(fileNotOpenedEvent.msg, path)
}

func (l *MyLogger) ConfigNotExportedLog(path string) {
	l.Error().Timestamp().Fields(defaultFields).Msgf(configNotExportedEvent.msg, path)
}

func (l *MyLogger) InvalidArgRangeLog(arg string) {
	l.Warn().Timestamp().Fields(defaultFields).Msgf(invalidArgRangeEvent.msg, arg)
}

func (l *MyLogger) EnvVarMissingLog(envVar string) {
	l.Warn().Timestamp().Fields(defaultFields).Msgf(missingEnvVarEvent.msg, envVar)
}

func (l *MyLogger) ConfigInfoLog(ident, value string) {
	l.Info().Timestamp().Fields(defaultFields).Msgf(configInfoEvent.msg, ident, value)
}
