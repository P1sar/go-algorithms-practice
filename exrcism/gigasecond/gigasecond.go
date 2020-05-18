package gigasecond

import (
	"math"
	"time"
)

func AddGigasecond(t time.Time) time.Time {
	gigananosecond_after := t.UnixNano() + int64(math.Pow10(18))
	time_after := time.Unix(0, gigananosecond_after)
	return  time_after
}
