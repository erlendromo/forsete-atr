package util

import "fmt"

type LogData struct {
	start      string
	zone       string
	endpoint   string
	method     string
	hasPayload bool
	status     int
	took       int64
}

func NewLogData(start, zone, endpoint, method string, hasPayload bool, status int, took int64) *LogData {
	return &LogData{
		start:      start,
		zone:       zone,
		endpoint:   endpoint,
		method:     method,
		hasPayload: hasPayload,
		status:     status,
		took:       took,
	}
}

func (l *LogData) PrintLog() {
	fmt.Printf(
		"\nRequest recieved: %s (%s)\nEndpoint: %s\nMethod: %s\nHas payload: %v\nStatus: %d\nTook: %dms\n\n",
		l.start,
		l.zone,
		l.endpoint,
		l.method,
		l.hasPayload,
		l.status,
		l.took,
	)
}
