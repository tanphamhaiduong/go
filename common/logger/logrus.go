package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type logrusLogEntry struct {
	entry *logrus.Entry
}

type logrusLogger struct {
	logger *logrus.Logger
}

func getFormatter(isJSON bool) logrus.Formatter {
	if isJSON {
		return &logrus.JSONFormatter{}
	}
	return &logrus.TextFormatter{
		FullTimestamp:          true,
		DisableLevelTruncation: true,
	}
}

func newLogrusLogger(config Configuration) (Logger, error) {
	logLevel := config.ConsoleLevel
	if logLevel == "" {
		logLevel = config.FileLevel
	}

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return nil, err
	}

	stdOutHandler := os.Stdout
	fileHandler := &lumberjack.Logger{
		Filename: config.FileLocation,
		MaxSize:  100,
		Compress: true,
		MaxAge:   28,
	}
	lLogger := &logrus.Logger{
		Out:       stdOutHandler,
		Formatter: getFormatter(config.ConsoleJSONFormat),
		Hooks:     make(logrus.LevelHooks),
		Level:     level,
	}

	if config.EnableConsole && config.EnableFile {
		lLogger.SetOutput(io.MultiWriter(stdOutHandler, fileHandler))
	} else {
		if config.EnableFile {
			lLogger.SetOutput(fileHandler)
			lLogger.SetFormatter(getFormatter(config.FileJSONFormat))
		}
	}

	return &logrusLogger{
		logger: lLogger,
	}, nil
}

// Debugf ...
func (l *logrusLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

// Infof ...
func (l *logrusLogger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

// Warnf ...
func (l *logrusLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

// Errorf ...
func (l *logrusLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

// Fatalf ...
func (l *logrusLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

// Panicf ...
func (l *logrusLogger) Panicf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

// WithFields ...
func (l *logrusLogger) WithFields(fields Fields) Logger {
	return &logrusLogEntry{
		entry: l.logger.WithFields(convertToLogrusFields(fields)),
	}
}

// Debugf ...
func (l *logrusLogEntry) Debugf(format string, args ...interface{}) {
	l.Debugf(format, args...)
}

// Infof ...
func (l *logrusLogEntry) Infof(format string, args ...interface{}) {
	l.Infof(format, args...)
}

// Warnf ...
func (l *logrusLogEntry) Warnf(format string, args ...interface{}) {
	l.Warnf(format, args...)
}

// Errorf ...
func (l *logrusLogEntry) Errorf(format string, args ...interface{}) {
	l.Errorf(format, args...)
}

// Fatalf ...
func (l *logrusLogEntry) Fatalf(format string, args ...interface{}) {
	l.Fatalf(format, args...)
}

// Panicf ...
func (l *logrusLogEntry) Panicf(format string, args ...interface{}) {
	l.Fatalf(format, args...)
}

// WithFields ...
func (l *logrusLogEntry) WithFields(fields Fields) Logger {
	return l.WithFields(fields)
}

func convertToLogrusFields(fields Fields) logrus.Fields {
	logrusFields := logrus.Fields{}
	for index, val := range fields {
		logrusFields[index] = val
	}
	return logrusFields
}
