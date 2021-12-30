package customLog

import (
	"github.com/sirupsen/logrus"
)

func InfoLn(args ...interface{}) {
	file, line, funName := trace()
	logrus.WithFields(logrus.Fields{
		"[FILE]": file,
		"[FUNC NAME]": funName,
		"[LINE]": line,
	}).Infoln(args...)
}

func InfoF(format string, args ...interface{}) {
	file, line, funName := trace()
	logrus.WithFields(logrus.Fields{
		"[FILE]": file,
		"[FUNC NAME]": funName,
		"[LINE]": line,
	}).Infof(format, args...)
}

func Warn(args ...interface{}) {
	file, line, funName := trace()
	logrus.WithFields(logrus.Fields{
		"[FILE]": file,
		"[FUNC NAME]": funName,
		"[LINE]": line,
	}).Warn(args...)
}

func WarnLn(args ...interface{}) {
	file, line, funName := trace()
	logrus.WithFields(logrus.Fields{
		"[FILE]": file,
		"[FUNC NAME]": funName,
		"[LINE]": line,
	}).Warnln(args...)
}

func WarnF(format string, args ...interface{}) {
	file, line, funName := trace()
	logrus.WithFields(logrus.Fields{
		"[FILE]": file,
		"[FUNC NAME]": funName,
		"[LINE]": line,
	}).Warnf(format, args...)
}

func Error(args ...interface{}) {
	file, line, funName := trace()
	logrus.WithFields(logrus.Fields{
		"[FILE]": file,
		"[FUNC NAME]": funName,
		"[LINE]": line,
	}).Error(args...)
}

func ErrorLn(args ...interface{}) {
	file, line, funName := trace()
	logrus.WithFields(logrus.Fields{
		"[FILE]": file,
		"[FUNC NAME]": funName,
		"[LINE]": line,
	}).Errorln(args...)
}

func ErrorF(format string, args ...interface{}) {
	file, line, funName := trace()
	logrus.WithFields(logrus.Fields{
		"[FILE]": file,
		"[FUNC NAME]": funName,
		"[LINE]": line,
	}).Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	file, line, funName := trace()
	logrus.WithFields(logrus.Fields{
		"[FILE]": file,
		"[FUNC NAME]": funName,
		"[LINE]": line,
	}).Fatal(args...)
}

func FatalLn(args ...interface{}) {
	file, line, funName := trace()
	logrus.WithFields(logrus.Fields{
		"[FILE]": file,
		"[FUNC NAME]": funName,
		"[LINE]": line,
	}).Fatalln(args...)
}

func FatalF(format string, args ...interface{}) {
	file, line, funName := trace()
	logrus.WithFields(logrus.Fields{
		"[FILE]": file,
		"[FUNC NAME]": funName,
		"[LINE]": line,
	}).Fatalf(format, args...)
}
