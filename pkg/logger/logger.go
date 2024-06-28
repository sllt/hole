package logger

import (
	"github.com/sllt/log"
	"os"
)

func New() *logger {

	l := log.NewWithOptions(os.Stderr, log.Options{ReportTimestamp: true})
	l.SetLevel(log.DebugLevel)

	return &logger{
		logger: l,
	}

}

type logger struct {
	logger *log.Logger
}

func (l *logger) SetLevel(lvl int) {
	l.logger.SetLevel(log.Level(lvl))
}
func (l *logger) Debug(format string, v ...interface{}) {
	l.logger.Debugf(format, v...)
}
func (l *logger) Info(format string, v ...interface{}) {
	l.logger.Infof(format, v...)
}
func (l *logger) Warn(format string, v ...interface{}) {
	l.logger.Warnf(format, v...)
}

func (l *logger) Error(format string, v ...interface{}) {
	l.logger.Errorf(format, v...)
}
