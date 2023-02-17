package utils

import (
	"fmt"
	"time"

	"github.com/darabuchi/utils/xtime"
)

func FormatTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	if t.Minute() == 0 && t.Hour() == 0 {
		return t.Format("2006年01月02日")
	}

	return t.Format("2006年01月02日 15:04")
}

func FormatDuration(d time.Duration) string {
	if d >= xtime.Year {
		years := d.Hours() / 24 / 365
		d = d - time.Duration(years)*xtime.Year

		mouths := d.Hours() / 24 / 30
		d = d - time.Duration(mouths)*xtime.Month

		days := d.Hours() / 24
		d = d - time.Duration(days)*xtime.Day

		hours := d.Hours()
		d = d - time.Duration(hours)*xtime.Hour

		minutes := d.Minutes()
		d = d - time.Duration(minutes)*xtime.Minute

		seconds := d.Seconds()

		if seconds == 0 {
			if minutes == 0 {
				if hours == 0 {
					if days == 0 {
						if mouths == 0 {
							return fmt.Sprintf("%d年", int(years))
						}
						return fmt.Sprintf("%d年%d月", int(years), int(mouths))
					}
					return fmt.Sprintf("%d年%d月%d天", int(years), int(mouths), int(days))
				}
				return fmt.Sprintf("%d年%d月%d天%d小时", int(years), int(mouths), int(days), int(hours))
			}
			return fmt.Sprintf("%d年%d月%d天%d小时%d分钟", int(years), int(mouths), int(days), int(hours), int(minutes))
		}
		return fmt.Sprintf("%d年%d月%d天%d小时%d分钟%d秒", int(years), int(mouths), int(days), int(hours), int(minutes), int(seconds))
	}

	if d >= xtime.Month {
		mouths := d.Hours() / 24 / 30
		d = d - time.Duration(mouths)*xtime.Month

		days := d.Hours() / 24
		d = d - time.Duration(days)*xtime.Day

		hours := d.Hours()
		d = d - time.Duration(hours)*xtime.Hour

		minutes := d.Minutes()
		d = d - time.Duration(minutes)*xtime.Minute

		seconds := d.Seconds()

		if seconds == 0 {
			if minutes == 0 {
				if hours == 0 {
					if days == 0 {
						return fmt.Sprintf("%d月", int(mouths))
					}
					return fmt.Sprintf("%d月%d天", int(mouths), int(days))
				}
				return fmt.Sprintf("%d月%d天%d小时", int(mouths), int(days), int(hours))
			}
			return fmt.Sprintf("%d月%d天%d小时%d分钟", int(mouths), int(days), int(hours), int(minutes))
		}
		return fmt.Sprintf("%d月%d天%d小时%d分钟%d秒", int(mouths), int(days), int(hours), int(minutes), int(seconds))
	}

	if d >= xtime.Day {
		days := d.Hours() / 24
		d = d - time.Duration(days)*xtime.Day

		hours := d.Hours()
		d = d - time.Duration(hours)*xtime.Hour

		minutes := d.Minutes()
		d = d - time.Duration(minutes)*xtime.Minute

		seconds := d.Seconds()

		if seconds == 0 {
			if minutes == 0 {
				if hours == 0 {
					return fmt.Sprintf("%d天", int(days))
				}
				return fmt.Sprintf("%d天%d小时", int(days), int(hours))
			}
			return fmt.Sprintf("%d天%d小时%d分钟", int(days), int(hours), int(minutes))
		}
		return fmt.Sprintf("%d天%d小时%d分钟%d秒", int(days), int(hours), int(minutes), int(seconds))
	}

	if d >= xtime.Hour {
		hours := d.Hours()
		d = d - time.Duration(hours)*xtime.Hour

		minutes := d.Minutes()
		d = d - time.Duration(minutes)*xtime.Minute

		seconds := d.Seconds()

		if seconds == 0 {
			if minutes == 0 {
				return fmt.Sprintf("%d小时", int(hours))
			}
			return fmt.Sprintf("%d小时%d分钟", int(hours), int(minutes))
		}
		return fmt.Sprintf("%d小时%d分钟%d秒", int(hours), int(minutes), int(seconds))
	}

	if d >= xtime.Minute {
		minutes := d.Minutes()
		d = d - time.Duration(minutes)*xtime.Minute

		seconds := d.Seconds()
		d = d - time.Duration(seconds)*xtime.Second

		milliseconds := d.Milliseconds()

		if milliseconds == 0 {
			if seconds == 0 {
				return fmt.Sprintf("%d分钟", int(minutes))
			}
			return fmt.Sprintf("%d分钟%d秒", int(minutes), int(seconds))
		}

		return fmt.Sprintf("%d分钟%d秒%d毫秒", int(minutes), int(seconds), int(milliseconds))
	}

	if d >= xtime.Second {
		seconds := d.Seconds()
		d = d - time.Duration(seconds)*xtime.Second

		milliseconds := d.Milliseconds()
		d = d - time.Duration(milliseconds)*xtime.Millisecond

		microseconds := d.Microseconds()

		if microseconds == 0 {
			if milliseconds == 0 {
				return fmt.Sprintf("%d秒", int(seconds))
			}
			return fmt.Sprintf("%d秒%d毫秒", int(seconds), int(milliseconds))
		}
		return fmt.Sprintf("%d秒%d毫秒%d微秒", int(seconds), int(milliseconds), int(microseconds))
	}

	if d >= xtime.Millisecond {
		milliseconds := d.Milliseconds()
		d = d - time.Duration(milliseconds)*xtime.Millisecond

		microseconds := d.Microseconds()
		d = d - time.Duration(microseconds)*xtime.Microsecond

		nanoseconds := d.Nanoseconds()

		if nanoseconds == 0 {
			if microseconds == 0 {
				return fmt.Sprintf("%d毫秒", int(milliseconds))
			}
			return fmt.Sprintf("%d毫秒%d微秒", int(milliseconds), int(microseconds))
		}
	}

	if d >= xtime.Microsecond {
		microseconds := d.Microseconds()
		d = d - time.Duration(microseconds)*xtime.Microsecond

		nanoseconds := d.Nanoseconds()

		if nanoseconds == 0 {
			return fmt.Sprintf("%d微秒", int(microseconds))
		}
		return fmt.Sprintf("%d微秒%d纳秒", int(microseconds), int(nanoseconds))
	}

	if d >= xtime.Nanosecond {
		return fmt.Sprintf("%d纳秒", d.Nanoseconds())
	}

	return "0s"
}
