package util

import (
	"fmt"
	"time"
)

var start time.Time

func StartUTCTimer() {
	start = time.Now().UTC()
}

func UpTimeInSeconds() string {
	return fmt.Sprintf("%0.fs", time.Since(start).Seconds())
}

func UpTimeInHHMMSS() string {
	return time.Since(start).String()
}
