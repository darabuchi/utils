package log

import "github.com/sirupsen/logrus"

func Trace(args ...interface{}) {
	std.Trace(args...)
}

func Debug(args ...interface{}) {
	std.Debug(args...)
}

func Print(args ...interface{}) {
	std.Print(args...)
}

func Info(args ...interface{}) {
	std.Info(args...)
}

func Warn(args ...interface{}) {
	std.Warn(args...)
}

func Warning(args ...interface{}) {
	std.Warning(args...)
}

func Error(args ...interface{}) {
	std.Error(args...)
}

func Panic(args ...interface{}) {
	std.Panic(args...)
}

func Fatal(args ...interface{}) {
	std.Fatal(args...)
}

func TraceFn(fn logrus.LogFunction) {
	std.TraceFn(fn)
}

func DebugFn(fn logrus.LogFunction) {
	std.DebugFn(fn)
}

func PrintFn(fn logrus.LogFunction) {
	std.PrintFn(fn)
}

func InfoFn(fn logrus.LogFunction) {
	std.InfoFn(fn)
}

func WarnFn(fn logrus.LogFunction) {
	std.WarnFn(fn)
}

func WarningFn(fn logrus.LogFunction) {
	std.WarningFn(fn)
}

func ErrorFn(fn logrus.LogFunction) {
	std.ErrorFn(fn)
}

func PanicFn(fn logrus.LogFunction) {
	std.PanicFn(fn)
}

func FatalFn(fn logrus.LogFunction) {
	std.FatalFn(fn)
}

func Tracef(format string, args ...interface{}) {
	std.Tracef(format, args...)
}

func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}

func Printf(format string, args ...interface{}) {
	std.Printf(format, args...)
}

func Infof(format string, args ...interface{}) {
	std.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	std.Warnf(format, args...)
}

func Warningf(format string, args ...interface{}) {
	std.Warningf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	std.Panicf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	std.Fatalf(format, args...)
}

func Traceln(args ...interface{}) {
	std.Traceln(args...)
}

func Debugln(args ...interface{}) {
	std.Debugln(args...)
}

func Println(args ...interface{}) {
	std.Println(args...)
}

func Infoln(args ...interface{}) {
	std.Infoln(args...)
}

func Warnln(args ...interface{}) {
	std.Warnln(args...)
}

func Warningln(args ...interface{}) {
	std.Warningln(args...)
}

func Errorln(args ...interface{}) {
	std.Errorln(args...)
}

func Panicln(args ...interface{}) {
	std.Panicln(args...)
}

func Fatalln(args ...interface{}) {
	std.Fatalln(args...)
}
