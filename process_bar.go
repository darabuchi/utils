package utils

import (
	"fmt"
	"strings"
)

func ProcessBarAscii(total, done int64, wight int64) string {
	const (
		doneChar = "▓"
		waitChar = "░"
	)

	if total <= done {
		return fmt.Sprintf("|%s|100%%[%d/%d]", strings.Repeat(doneChar, int(wight)), total, total)
	} else if done == 0 {
		return fmt.Sprintf("|%s|0%%(0/%d)", strings.Repeat(waitChar, int(wight)), total)
	}

	a := make([]string, int(wight), int(wight))
	p := float64(done) * 100 / float64(total)
	pt := int64(p) * wight / 100

	for i := int64(0); i < wight; i++ {
		// if i < pt {
		// 	a[i] = doneChar
		// } else if i == pt {
		// 	a[i] = fmt.Sprintf("%.2f%%", p)
		// } else {
		// 	a[i] = waitChar
		// }

		if i <= pt {
			a[i] = doneChar
		} else {
			a[i] = waitChar
		}
	}

	return fmt.Sprintf("|%s|%.2f%%[%d/%d]", strings.Join(a, ""), p, done, total)
}
