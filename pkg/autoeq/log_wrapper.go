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

var loggerInstance *MyLogger
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
	var writer = os.Stdout
	f, err := prepareLogFile()

	if err == nil {
		writer = f
	}

	logger := zerolog.New(writer).With().Timestamp().Fields(defaultFields).Logger()
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

func GetLogger() *MyLogger {
	if loggerInstance == nil {
		loggerInstance = NewLogger()
	}

	return loggerInstance
}

func (l *MyLogger) SuccessLog(context string) {
	l.Info().Msgf(successEvent.msg, context)
}

func (l *MyLogger) FileNotOpenedLog(path string) {
	l.Error().Msgf(fileNotOpenedEvent.msg, path)
}

func (l *MyLogger) ConfigNotExportedLog(path string) {
	l.Error().Msgf(configNotExportedEvent.msg, path)
}

func (l *MyLogger) InvalidArgRangeLog(arg string) {
	l.Warn().Msgf(invalidArgRangeEvent.msg, arg)
}

func (l *MyLogger) EnvVarMissingLog(envVar string) {
	l.Warn().Msgf(missingEnvVarEvent.msg, envVar)
}

func (l *MyLogger) ConfigInfoLog(ident, value string) {
	l.Info().Msgf(configInfoEvent.msg, ident, value)
}
