package logutils

import (
	"encoding/json"
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
)

func LogSuccess(messages ...string) {
	if len(messages) != 0 {
		var sb strings.Builder
		sb.WriteString(ANSI_Green)
		for _, m := range messages {
			sb.WriteString(m)
		}
		sb.WriteString(ANSI_Reset)
		log.Println(sb.String())
	}
}

func LogInfo(messages ...string) {
	if len(messages) != 0 {
		var sb strings.Builder
		sb.WriteString(ANSI_Yellow)
		for _, m := range messages {
			sb.WriteString(m)
		}
		sb.WriteString(ANSI_Reset)
		log.Println(sb.String())
	}
}

func LogError(e error) {
	var sb strings.Builder
	sb.WriteString(ANSI_Red)
	sb.WriteString(e.Error())
	sb.WriteString(ANSI_Reset)
	log.Println(sb.String())
}

func LogAsJSON(v interface{}) {
	bytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Printf("%s failed to marshal JSON: %v %s\n", ANSI_Red, err, ANSI_Reset)
		return
	}
	var sb strings.Builder
	sb.WriteString(ANSI_Green)
	sb.WriteString(string(bytes))
	sb.WriteString(ANSI_Reset)
	log.Println(sb.String())
}
