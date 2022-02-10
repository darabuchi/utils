package log

import (
	"bytes"
	"fmt"
	"github.com/petermattis/goid"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
	"strings"
)

type Formatter struct {
	module string
}

const (
	colorBlack   = "\u001B[30m"
	colorRed     = "\u001B[31m"
	colorGreen   = "\u001B[32m"
	colorYellow  = "\u001B[33m"
	colorBlue    = "\u001B[34m"
	colorMagenta = "\u001B[35m"
	colorCyan    = "\u001B[36m"
	colorGray    = "\u001B[37m"
	colorWhite   = "\u001B[38m"
)

const (
	endColor = "\u001B[0m"
)

func (p *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b bytes.Buffer

	if p.module != "" {
		b.WriteString(p.module)
		b.WriteString(" ")
	}

	b.WriteString(fmt.Sprintf("(%d.%d) ", pid, goid.Get()))

	b.WriteString(entry.Time.Format("2006-01-02 15:04:05.9999Z07:00"))

	color := p.getColorByLevel(entry.Level)

	b.WriteString(color)
	b.WriteString(" [")
	b.WriteString(entry.Level.String()[:4])
	b.WriteString("] ")
	b.WriteString(endColor)

	b.WriteString(strings.TrimSpace(entry.Message))

	var callerName string
	pc, file, callerLine, ok := runtime.Caller(7)
	if ok {
		callerName = runtime.FuncForPC(pc).Name()
	}

	b.WriteString(color)
	b.WriteString(" (")
	b.WriteString(path.Join(p.getPackageName(callerName), path.Base(file)))
	b.WriteString(":")
	b.WriteString(fmt.Sprintf("%d", callerLine))
	b.WriteString(")")
	b.WriteString(endColor)

	b.WriteByte('\n')

	return b.Bytes(), nil
}

func (p *Formatter) getColorByLevel(level logrus.Level) string {
	switch level {
	case logrus.DebugLevel, logrus.TraceLevel:
		return colorGreen

	case logrus.WarnLevel:
		return colorYellow

	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		return colorRed

	default:
		return colorGreen
	}
}

func (p *Formatter) getPackageName(f string) string {
	for {
		lastPeriod := strings.LastIndex(f, ".")
		lastSlash := strings.LastIndex(f, "/")
		if lastPeriod > lastSlash {
			f = f[:lastPeriod]
		} else {
			break
		}
	}

	return f
}

func (p *Formatter) SetModule(module string) {
	p.module = module
}
