package util

import (
	"fmt"
)

var logColors = map[string]string{
	"INFO":         CYAN,
	"SUCCESS":      GREEN,
	"CLIENT ERROR": YELLOW,
	"SERVER ERROR": RED,
	"MISC":         PURPLE,
}

func getTypeAndColor(logType string) (string, string) {
	color, found := logColors[logType]
	if !found {
		logType = "MISC"
		color = logColors[logType]
	}

	return logType, color
}

type Log interface {
	PrintLog(logType string)
}

type ReceiveLog struct {
	start    string
	zone     string
	client   string
	endpoint string
	method   string
}

func NewReceiveLog(start, zone, client, endpoint, method string) *ReceiveLog {
	return &ReceiveLog{
		start:    start,
		zone:     zone,
		client:   client,
		endpoint: endpoint,
		method:   method,
	}
}

func (rl *ReceiveLog) PrintLog(logType string) {
	logType, color := getTypeAndColor(logType)

	fmt.Printf(
		"\n%s%s%s\nRequest recieved: %s (%s)\nClient: %s\nEndpoint: %s\nMethod: %s\n",
		color,
		logType,
		RESET,
		rl.start,
		rl.zone,
		rl.client,
		rl.endpoint,
		rl.method,
	)
}

type ResponseLog struct {
	status int
	took   int64
	unit   string
}

func NewResponseLog(status int, took int64, unit string) *ResponseLog {
	return &ResponseLog{
		status: status,
		took:   took,
		unit:   unit,
	}
}

func (rl *ResponseLog) PrintLog(logType string) {
	logType, color := getTypeAndColor(logType)

	fmt.Printf(
		"\n%s%s%s\nStatus: %d\nTook: %d%s\n",
		color,
		logType,
		RESET,
		rl.status,
		rl.took,
		rl.unit,
	)
}
