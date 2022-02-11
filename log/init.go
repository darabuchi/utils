package log

import (
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

var (
	//std = logrus.New()
	std = newLogger()
	//logFmt = &nested.Formatter{
	//	FieldsOrder: []string{
	//		logrus.FieldKeyTime, logrus.FieldKeyLevel,
	//		logrus.FieldKeyFunc, logrus.FieldKeyMsg,
	//	},
	//	CustomCallerFormatter: func(f *runtime.Frame) string {
	//		return fmt.Sprintf("(%s %s:%d)", f.Function, path.Base(f.File), f.Line)
	//	},
	//	TimestampFormat:  "2006-01-02 15:04:05.9999Z07:00",
	//	HideKeys:         true,
	//	NoColors:         false,
	//	NoFieldsColors:   false,
	//	NoFieldsSpace:    true,
	//	ShowFullLevel:    true,
	//	NoUppercaseLevel: true,
	//	TrimMessages:     true,
	//	CallerFirst:      true,
	//}

	logFmt = &Formatter{}

	pid = os.Getpid()
)

func init() {
	//std.SetFormatter(logFmt)
	//std.SetReportCaller(false)
	//std.SetOutput(os.Stdout)
}

func New() *Logger {
	return std
}

func SetOutput(write io.Writer) {
	std.SetOutput(zapcore.AddSync(write))
}

func SetLevel(level Level) {
	std.SetLevel(level)
}

//func SetModule(module string) {
//	logFmt.SetModule(module)
//}

func AddOutput(write io.Writer) {
	std.AddOutput(zapcore.AddSync(write))
}

func Sync() {
	std.Sync()
}
