package xtime

import "time"

const (
	Nanosecond  = time.Nanosecond
	Microsecond = time.Microsecond
	Millisecond = time.Millisecond
	Second      = time.Second
	Minute      = time.Minute
	Hour        = time.Hour
	Day         = time.Hour * 24
	Week        = Day * 7
	Month       = Day * 30
	Quarter     = Day * 91
	Year        = Day * 365
)
