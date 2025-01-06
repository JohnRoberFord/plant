package sign

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func SignString(input string, key string) (string, error) {
	h := hmac.New(sha256.New, []byte(key))
	_, err := h.Write([]byte(input))
	if err != nil {
		return "", err
	}
	hash := h.Sum(nil)
	return hex.EncodeToString(hash), nil
}

func IsValid(msg string, recieved string, keyString string) (bool, error) {
	mac1, err := hex.DecodeString(recieved)
	if err != nil {
		return false, err
	}
	mac2, err := SignString(msg, keyString)
	if err != nil {
		return false, err
	}

	return hmac.Equal(mac1, []byte(mac2)), nil
}
