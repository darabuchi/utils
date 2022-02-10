package log

import "github.com/sirupsen/logrus"

type Level logrus.Level

const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)
