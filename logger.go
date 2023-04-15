package lark

import (
	"github.com/sirupsen/logrus"
)

var logger = func() *logrus.Logger {
	l := logrus.New()
	l.SetReportCaller(true)
	l.SetLevel(logrus.InfoLevel)
	return l
}()
