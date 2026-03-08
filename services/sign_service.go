package services

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"strings"
)

func SignHMACSHA512(text, key string) (string, error) {

	if strings.TrimSpace(text) == "" {
		return "", &ValidationError{Field: "Text", Message: "text не может быть пустым"}
	}
	if strings.TrimSpace(key) == "" {
		return "", &ValidationError{Field: "Key", Message: "key не может быть пустым"}
	}

	hasher := hmac.New(sha512.New, []byte(key))

	hasher.Write([]byte(text))

	signature := hasher.Sum(nil)

	return hex.EncodeToString(signature), nil
}
