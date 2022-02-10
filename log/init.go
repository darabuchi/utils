package log

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

var (
	std    = logrus.New()
	logFmt = &nested.Formatter{
		FieldsOrder: []string{
			logrus.FieldKeyTime, logrus.FieldKeyLevel,
			logrus.FieldKeyFunc, logrus.FieldKeyMsg,
		},
		CustomCallerFormatter: func(f *runtime.Frame) string {
			return fmt.Sprintf("(%s %s:%d)", f.Function, path.Base(f.File), f.Line)
		},
		TimestampFormat:  "2006-01-02 15:04:05.9999Z07:00",
		HideKeys:         true,
		NoColors:         false,
		NoFieldsColors:   false,
		NoFieldsSpace:    true,
		ShowFullLevel:    true,
		NoUppercaseLevel: true,
		TrimMessages:     true,
		CallerFirst:      true,
	}
)

func init() {

	std.SetFormatter(logFmt)
	std.SetOutput(os.Stdout)
}

func New() *logrus.Logger {
	return std
}

func SetOutput(out io.Writer) {
	std.SetOutput(out)
}

func SetFormatter(formatter logrus.Formatter) {
	std.SetFormatter(formatter)
}

func SetReportCaller(include bool) {
	std.SetReportCaller(include)
}

func AddOutput(out io.Writer) {
	logrus.AddHook(lfshook.NewHook(lfshook.WriterMap{
		logrus.TraceLevel: out,
		logrus.DebugLevel: out,
		logrus.InfoLevel:  out,
		logrus.WarnLevel:  out,
		logrus.ErrorLevel: out,
		logrus.FatalLevel: out,
		logrus.PanicLevel: out,
	}, logFmt))
}
