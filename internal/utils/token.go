package utils

import "strings"

func GetTokenFromAuthorizationHeader(authorization string) string {
	if authorization == "" {
		return ""
	}

	token := ""
	if strings.Contains(authorization, " ") {
		token = strings.TrimPrefix(authorization, "Bearer ")
	}

	if token == "" {
		return ""
	}

	return token
}
