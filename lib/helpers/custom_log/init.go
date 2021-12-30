package customLog

import (
	"io/ioutil"
	"os"

	"github.com/alhamsya/boilerplate-go/domain/constants"
	"github.com/sirupsen/logrus"
)

func InitializeLogging(accessLogFile, errorLogFile string) error {
	accessLog, err := OpenFile(accessLogFile)
	if err != nil {
		return err
	}

	errorLog, err := OpenFile(errorLogFile)
	if err != nil {
		return err
	}

	logrus.SetOutput(ioutil.Discard)

	logrus.AddHook(&WriterHook{
		Writer: accessLog,
		LogLevels: []logrus.Level{
			logrus.InfoLevel,
			logrus.WarnLevel,
		},
	})

	logrus.AddHook(&WriterHook{
		Writer: errorLog,
		LogLevels: []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
		},
	})

	// WRITE TO CONSOLE
	logrus.AddHook(&WriterHook{
		Writer: os.Stdout,
		LogLevels: []logrus.Level{
			logrus.InfoLevel,
			logrus.WarnLevel,
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
			logrus.DebugLevel,
			logrus.TraceLevel,
		},
	})

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:            true,
		FullTimestamp:          true,
		TimestampFormat:        constCommon.DateTimeWithTimezone,
		DisableLevelTruncation: true,
		DisableSorting:         true,
	})

	return nil
}
