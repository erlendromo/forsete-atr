package util

/*
Endpoints
*/
const (
	SLASH         = "/"
	FORSETE_ATR   = "forsete-atr"
	VERSION       = "v1"
	BASE_ENDPOINT = SLASH + FORSETE_ATR + SLASH + VERSION + SLASH

	// Swaggo
	SWAGGO               = "swaggo"
	DOC_JSON             = "doc.json"
	SWAGGO_ENDPOINT      = BASE_ENDPOINT + SWAGGO + SLASH
	SWAGGO_DOCS_ENDPOINT = SWAGGO_ENDPOINT + DOC_JSON

	// Status
	STATUS          = "status"
	STATUS_ENDPOINT = BASE_ENDPOINT + STATUS + SLASH

	// Models
	MODELS                       = "models"
	REGION_SEGMENTATION_MODELS   = "region-segmentation-models"
	LINE_SEGMENTATION_MODELS     = "line-segmentation-models"
	TEXT_RECOGNITION_MODELS      = "text-recognition-models"
	MODELS_ENDPOINT              = BASE_ENDPOINT + MODELS + SLASH
	REGION_SEGMENTATION_ENDPOINT = MODELS_ENDPOINT + REGION_SEGMENTATION_MODELS + SLASH
	LINE_SEGMENTATION_ENDPOINT   = MODELS_ENDPOINT + LINE_SEGMENTATION_MODELS + SLASH
	TEXT_RECOGNITION_ENDPOINT    = MODELS_ENDPOINT + TEXT_RECOGNITION_MODELS + SLASH

	// ATR
	ATR                        = "atr"
	BASIC_DOCUMENTS            = "basic-documents"
	TIPNOTE_DOCUMENTS          = "tipnote-documents"
	ATR_ENDPOINT               = BASE_ENDPOINT + ATR + SLASH
	BASIC_DOCUMENTS_ENDPOINT   = ATR_ENDPOINT + BASIC_DOCUMENTS + SLASH
	TIPNOTE_DOCUMENTS_ENDPOINT = ATR_ENDPOINT + TIPNOTE_DOCUMENTS + SLASH
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

/*
Color
*/
const (
	RESET  = "\033[0m"
	RED    = "\033[31m"
	GREEN  = "\033[32m"
	YELLOW = "\033[33m"
	PURPLE = "\033[35m"
	CYAN   = "\033[36m"
)

/*
Misc
*/
const (
	ASSETS              = "assets"
	REGION_SEGMENTATION = "regionsegmentation"
	LINE_SEGMENTATION   = "linesegmentation"
	TEXT_RECOGNITION    = "textrecognition"
)

const (
	INTERNAL_SERVER_ERROR = "internal server error"
)
