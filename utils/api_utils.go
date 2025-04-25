package utils

import (
	"log"
	"strings"
)

const (
	ANSI_Reset  = "\033[0m"
	ANSI_Black  = "\033[30m"
	ANSI_Red    = "\033[31m"
	ANSI_Green  = "\033[32m"
	ANSI_Yellow = "\033[33m"
	ANSI_Blue   = "\033[34m"
	ANSI_Purple = "\033[35m"
	ANSI_Cyan   = "\033[36m"
	ANSI_White  = "\033[37m"

	SuccessCode         = 0
	InternalErrorCode   = 1
	DBErrorCode         = 2
	DBConflictErrorCode = 4
	AuthorizedErrorCode = 8
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

func LogSuc(msg string) {
	var sb strings.Builder
	sb.WriteString(ANSI_Green)
	sb.WriteString(msg)
	sb.WriteString(ANSI_Reset)
	log.Println(sb.String())
}

func LogInfo(msg string) {
	var sb strings.Builder
	sb.WriteString(ANSI_Yellow)
	sb.WriteString(msg)
	sb.WriteString(ANSI_Reset)
	log.Println(sb.String())
}

func LogError(e error) {
	var sb strings.Builder
	sb.WriteString(ANSI_Red)
	sb.WriteString(e.Error())
	sb.WriteString(ANSI_Reset)
	log.Println(sb.String())
}
