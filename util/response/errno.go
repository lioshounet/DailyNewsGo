package response

const (
	Success    = 0
	ParamError = 1
	UNKNOWN    = 99
)

var MsgEN = map[int]string{
	Success:    "success",
	ParamError: "param error",
	UNKNOWN:    "unknown error",
}

func GetMsg(code int) string {
	if msg, ok := MsgEN[code]; ok {
		return msg
	}
	return MsgEN[UNKNOWN]
}
