package utils

const (
	SuccessCode     = 0
	ParserErrorCode = 1
	DBErrorCode     = 2
	OtherErrorCode  = 99

	Success = "success"
	Error   = "error"
)

type APIResponse struct {
	Code    int         `json:"code"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseOK(msg string, data any) APIResponse {
	return APIResponse{
		Code:    SuccessCode,
		Error:   "",
		Message: msg,
		Data:    data,
	}
}

func ResponseError(code int, errMsg string) APIResponse {
	return APIResponse{
		Code:    code,
		Error:   errMsg,
		Message: "",
	}
}

// Response template
// func RT(code int, msg any, obj any) (m map[string]any) {
// 	m = map[string]any{
// 		"code": code,
// 	}
// 	switch code {
// 	case ResponseSuccessCode:
// 		if msg != nil {
// 			m["message"] = msg
// 		} else {
// 			m["message"] = "Success"
// 		}
// 	case ResponseFailed:
// 		m["error"] = msg
// 	}
// 	if obj != nil {
// 		m["data"] = obj
// 	}
// 	return
// }
