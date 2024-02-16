package middleware

import (
	"strings"
)

func GetToken(token string) string {
	if strings.Contains(token, "Token ") {
		return strings.Split(token, "Token ")[1]
	}
	return ""
}
