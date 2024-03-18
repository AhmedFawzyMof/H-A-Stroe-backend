package middleware

import (
	"net/http"
	"strings"
)

func SetCors(res http.ResponseWriter, method string) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", method+", OPTIONS")
	res.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func GetToken(token string) string {
	if strings.Contains(token, "Token ") {
		return strings.Split(token, "Token ")[1]
	}
	return ""
}
