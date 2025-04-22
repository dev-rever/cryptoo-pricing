package utils

import (
	"fmt"
	"log"
	"strings"
)

const (
	ANSI_Red   = "\033[31m"
	ANSI_Green = "\033[32m"
	ANSI_Reset = "\033[0m"
)

type Logger struct {
	Code int
	Msg  string
}

func Log(code int, msg string) (logger Logger) {
	var sb strings.Builder
	if code > SuccessCode {
		sb.WriteString(ANSI_Red)
	} else {
		sb.WriteString(ANSI_Green)
	}
	sb.WriteString(fmt.Sprintf("code: %d, msg: %s%v", code, msg, ANSI_Reset))
	log.Println(sb.String())
	logger = Logger{
		Code: code,
		Msg:  msg,
	}
	return
}

func (l *Logger) Print() {
	log.Printf("code: %d, msg: %s", l.Code, l.Msg)
}
