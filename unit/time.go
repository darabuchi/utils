package unit

import (
	"fmt"
	"time"

	"github.com/darabuchi/utils/language"
	"github.com/darabuchi/utils/xtime"
)

func Duration2Str(t time.Duration) string {
	return newDuration(t).String()
}

func Duration2StrWithLanguage(t time.Duration, languageCode language.LanguageCode) string {
	return newDuration(t).StringWithLanguage(languageCode)
}

type duration struct {
	day, hour, min, second, millisecond, duration time.Duration
}

func newDuration(t time.Duration) *duration {
	p := &duration{}

	if t > xtime.Day {
		p.day = t / xtime.Day
	}
	t = t - p.day*xtime.Day

	if t > xtime.Hour {
		p.hour = t / xtime.Hour
	}
	t = t - p.hour*xtime.Hour

	if t > xtime.Minute {
		p.min = t / xtime.Minute
	}
	t = t - p.min*xtime.Minute

	if t > xtime.Second {
		p.second = t / xtime.Second
	}
	t = t - p.second*xtime.Second

	if t > xtime.Millisecond {
		p.millisecond = t / xtime.Millisecond
	}
	t = t - p.millisecond*xtime.Millisecond

	p.duration = t

	return p
}

func (p *duration) String() string {
	return fmt.Sprintf("%dd%dh%dm%ds%sms", p.day, p.hour, p.min, p.second, p.millisecond)
}

func (p *duration) StringWithLanguage(code language.LanguageCode) string {
	return fmt.Sprintf("%d%s%d%s%d%s%d%s",
		p.day, GetUnitCode2Show(Day, code),
		p.hour, GetUnitCode2Show(Hour, code),
		p.min, GetUnitCode2Show(Minute, code),
		p.second, GetUnitCode2Show(Second, code),
	)
}
