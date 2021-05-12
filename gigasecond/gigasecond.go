package gigasecond

import (
	"time"
)


func AddGigasecond(t time.Time) time.Time {
	//1e9=1.000.000.000 second
	return t.Add(time.Second * 1e9)
}

