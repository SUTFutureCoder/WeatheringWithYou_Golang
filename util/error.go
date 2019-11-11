package util

//直接从PHP LIB里面复制过来的 意外好用
// 通用错误码，占用100开头
const OK                      = 0;
const ERROR_PARAM_ERROR       = 100;
const ERROR_NETWORK_ERROR     = 101;
const ERROR_USER_NOT_LOGIN    = 102;
const ERROR_JSON_FORMAT_ERROR = 103;
const ERROR_INPUT_NULL        = 104;
const ERROR_RAL               = 105;
const ERROR_JSON_DECODE       = 106;
const ERROR_GEN_UUID          = 107;
const ERROR_UUID_MAX          = 108;
const ERROR_REDIS             = 109;
const ERROR_FUNC_NON_EXISTS   = 110;
const ERROR_IP_UNAUTHORIZED   = 111;
const ERROR_USER_UNAUTHORIZED = 112;

//DB
const ERROR_DB_CONNECT        = 206;
const ERROR_DB_INSERT         = 207;
const ERROR_DB_UPDATE         = 208;
const ERROR_DB_DELETE         = 209;
const ERROR_DB_SELECT         = 210;

const ERROR_NO_ERROR          = 301;

func GetErrorMessage (errorNo int) string{
	var ERROR_MESSAGE = map[int]string{
		ERROR_PARAM_ERROR:          "param error",
		ERROR_NETWORK_ERROR:        "network error",
		ERROR_USER_NOT_LOGIN:       "user not login",
		ERROR_JSON_FORMAT_ERROR:    "json format error",
		ERROR_INPUT_NULL:           "file get contents php input null",
		ERROR_RAL:                  "do http query by ral failed",
		ERROR_JSON_DECODE:          "decode json error",
		ERROR_GEN_UUID:             "gen uuid error",
		ERROR_UUID_MAX:             "uuid too long error",
		ERROR_REDIS:                "redis error",
		ERROR_FUNC_NON_EXISTS:      "function non exists error",
		ERROR_IP_UNAUTHORIZED:      "unauthorized IP",
		ERROR_USER_UNAUTHORIZED:	"unauthorized user",

		ERROR_DB_CONNECT:           "db connect error",
		ERROR_DB_INSERT:            "db insert error",
		ERROR_DB_UPDATE:            "db update error",
		ERROR_DB_DELETE:            "db delete error",
		ERROR_DB_SELECT:            "db select error",

		ERROR_NO_ERROR:             "error number error",
	}

	if ERROR_MESSAGE[errorNo] == "" {
		return ERROR_MESSAGE[ERROR_NO_ERROR]
	}
	return ERROR_MESSAGE[errorNo]
}
