package log

import (
	"github.com/sirupsen/logrus"
)

var Logger = func() *logrus.Logger {
	l := logrus.New()
	l.SetReportCaller(true)
	l.SetLevel(logrus.InfoLevel)
	return l
}()
