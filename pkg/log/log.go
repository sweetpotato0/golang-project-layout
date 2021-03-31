package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.InfoLevel)
}

func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}
func Info(args ...interface{}) {
	logrus.Info(args...)
}

func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}
