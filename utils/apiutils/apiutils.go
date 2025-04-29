package apiutils

const (
	SuccessCode         = 0
	InternalErrorCode   = 1
	DBErrorCode         = 2
	DBConflictErrorCode = 3
	AuthorizedErrorCode = 4
	UnknownErrorCode    = 99

	SuccessText = "success"
	ErrorText   = "error"
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
