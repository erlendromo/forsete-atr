package util

/*
Endpoints
*/
const (
	// Shared
	SLASH       = "/"
	FORSETE_ATR = "forsete-atr"
	VERSION     = "v2"
	DATA        = "data"

	BASE_ENDPOINT = SLASH + FORSETE_ATR + SLASH + VERSION + SLASH

	// Swaggo
	SWAGGO   = "swaggo"
	DOC_JSON = "doc.json"

	SWAGGO_ENDPOINT      = BASE_ENDPOINT + SWAGGO + SLASH
	SWAGGO_DOCS_ENDPOINT = SWAGGO_ENDPOINT + DOC_JSON

	// Auth
	AUTH     = "auth"
	REGISTER = "register"
	LOGIN    = "login"
	LOGOUT   = "logout"
	REFRESH  = "refresh"

	BASE_AUTH_ENDPOINT = BASE_ENDPOINT + AUTH + SLASH
	REGISTER_ENDPOINT  = BASE_AUTH_ENDPOINT + REGISTER + SLASH
	LOGIN_ENDPOINT     = BASE_AUTH_ENDPOINT + LOGIN + SLASH
	LOGOUT_ENDPOINT    = BASE_AUTH_ENDPOINT + LOGOUT + SLASH
	REFRESH_ENDPOINT   = BASE_AUTH_ENDPOINT + REFRESH + SLASH

	// Images
	IMAGES            = "images"
	UPLOAD            = "upload"
	IMAGE_ID_WILDCARD = "{imageID}"

	IMAGES_ENDPOINT        = BASE_ENDPOINT + IMAGES + SLASH
	UPLOAD_IMAGES_ENDPOINT = IMAGES_ENDPOINT + UPLOAD + SLASH
	IMAGE_BY_ID_ENDPOINT   = IMAGES_ENDPOINT + IMAGE_ID_WILDCARD + SLASH
	IMAGE_DATA_ENDPOINT    = IMAGE_BY_ID_ENDPOINT + DATA + SLASH

	// Outputs
	OUTPUTS            = "outputs"
	OUTPUT_ID_WILDCARD = "{outputID}"

	OUTPUTS_ENDPOINT      = IMAGE_BY_ID_ENDPOINT + OUTPUTS + SLASH
	OUTPUT_BY_ID_ENDPOINT = OUTPUTS_ENDPOINT + OUTPUT_ID_WILDCARD + SLASH
	OUTPUT_DATA_ENDPOINT  = OUTPUT_BY_ID_ENDPOINT + DATA + SLASH

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
	ATR          = "atr"
	ATR_ENDPOINT = BASE_ENDPOINT + ATR + SLASH

	// Status
	STATUS = "status"

	STATUS_ENDPOINT = BASE_ENDPOINT + STATUS + SLASH
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
	/*
		IMAGE_PNG        = "image/png"
		IMAGE_JPG        = "image/jpg"
		IMAGE_JPEG       = "image/jpeg"
	*/
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
