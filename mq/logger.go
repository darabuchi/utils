package mq

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Output(maxdepth int, s string) error {
	return nil
}
