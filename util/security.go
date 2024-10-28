package utils

import "crypto/rand"

func GenerateSecureNonce(length int) string {

	var allowedChars string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	bytes := make([]byte, length)
	_, e := rand.Read(bytes)
	if e != nil {
		panic("Can't generate secure random token")
	}

	for i, rng := range bytes {
		bytes[i] = allowedChars[rng%byte(len(allowedChars))]
	}
	return string(bytes)
}
