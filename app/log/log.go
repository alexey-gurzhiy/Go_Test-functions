package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Этот интерфейс конечно можно использовать, но по факту он заточен под logrus.
// Заменить потом на другой логгер не получится, нужно будет менять интерфейс
// Вполне возможно сейчас имеет смысл указать в App поле logger *logrus.Logger вместо log.Logger
type Logger interface {
	Trace(args ...interface{})
	Tracef(format string, args ...interface{})

	Debug(args ...interface{})
	Debugf(format string, args ...interface{})

	Print(args ...interface{})
	Printf(format string, args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})

	Warning(args ...interface{})
	Warningf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	Panic(args ...interface{})
	Panicf(format string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})

	WithFields(fields logrus.Fields) *logrus.Entry
}

func NewLogger(loglevel string) *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		DisableQuote:  true,
		FullTimestamp: true,
	})
	logger.SetOutput(os.Stdout)

	logger.SetLevel(getLogLevel(loglevel))
	return logger
}

func getLogLevel(loglevel string) logrus.Level {
	switch loglevel {
	case "Trace":
		return logrus.TraceLevel
	case "Debug":
		return logrus.DebugLevel
	case "Info":
		return logrus.InfoLevel
	case "Error":
		return logrus.ErrorLevel
	case "Fatal":
		return logrus.FatalLevel
	case "Panic":
		return logrus.PanicLevel
	default:
		return logrus.WarnLevel
	}
}
