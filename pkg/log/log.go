package log

import (
	"os"

	"github.com/onrik/logrus/filename"
	"github.com/sirupsen/logrus"
)

// New function initialize logrus and return a new logger
// We use an abstraction so that our logs are consistent and if there's anything that needs change
// related to logs, we can just change here
func New() *logrus.Logger {
	// Filename is a hook for logrus that adds file name and line number to the log as well.
	// It's useful for indicating where the log was originated from
	filenameHook := filename.NewHook()

	log := &logrus.Logger{
		Hooks:     make(logrus.LevelHooks),
		Out:       os.Stdout,
		Formatter: &logrus.TextFormatter{},
		Level:     logrus.InfoLevel,
	}

	log.Hooks.Add(filenameHook)

	return log
}
