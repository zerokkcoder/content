package utils

import "fmt"

func GetAuthKey(sessionID string) string {
	authkey := fmt.Sprintf("session_auth:%s", sessionID)
	return authkey
}

func GetSessionKey(username string) string {
	sessionKey := fmt.Sprintf("session_id:%s", username)
	return sessionKey
}
