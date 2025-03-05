package util

/*
Misc
*/
const (
	API_PORT         = "API_PORT"
	DEFAULT_API_PORT = "8080"
)

/*
Endpoints
*/
const (
	SLASH       = "/"
	FORSETE_ATR = "forsete-atr"
	VERSION     = "v1"
	BASIC       = "basic"
	TIPNOTE     = "tipnote"

	BASE_ENDPOINT    = SLASH + FORSETE_ATR + SLASH + VERSION + SLASH
	BASIC_ENDPOINT   = BASE_ENDPOINT + BASIC + SLASH
	TIPNOTE_ENDPOINT = BASE_ENDPOINT + TIPNOTE + SLASH
)

/*
Time
*/
const (
	TIME_FORMAT = "02.01.2006 15:04:05"
)

/*
Json
*/
const (
	CONTENT_TYPE     = "Content-Type"
	APPLICATION_JSON = "application/json"
)
