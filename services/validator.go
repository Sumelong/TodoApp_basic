package services

import (
	"log/slog"
	"strings"
)

func IsString(str string) string {
	str = strings.Trim(str, " '' ")
	str = strings.ToLower(str)

	if len(str) == 0 {
		slog.Info("empty string")
		return ""
	}
	return str
}
